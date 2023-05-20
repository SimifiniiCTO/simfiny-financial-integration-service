#!/usr/bin/env bash

# Try to hit the ngrok API inside docker compose, if this succeeds then that means webhooks are enabled for plaid for
# local development.

if [[ ${CLOUD_MAGIC} == "magic" ]]; then
  echo "[wrapper] Cloud magic detected, will use GitPod/CodeSpaces URL for webhooks instead."

  export SIMFINY_PLAID_WEBHOOKS_DOMAIN="https://${SIMFINY_API_DOMAIN_NAME}";
  export SIMFINY_PLAID_WEBHOOKS_ENABLED="true";
else
  WEBHOOKS_DOMAIN=$(curl http://ngrok:4040/api/tunnels -s -m 0.1 | perl -pe '/\"public_url\":\"https:\/\/(\S*?)\",/g; print $1;' | cut -d "{" -f1);

  if [[ ! -z "${WEBHOOKS_DOMAIN}" ]]; then
    echo "[wrapper] ngrok detected, webhooks should target: ${WEBHOOKS_DOMAIN}";

    # If the domain name has been derived then enable webhooks for plaid.
    echo "[wrapper] Plaid webhooks have been enabled...";
    export SIMFINY_PLAID_WEBHOOKS_DOMAIN=${WEBHOOKS_DOMAIN};
    export SIMFINY_PLAID_WEBHOOKS_ENABLED="true";
  else
    echo "[wrapper] ngrok not detected, webhooks will not be available..."
  fi
fi


# If the stripe API key, webhook secret and price ID are provided then enable billing for local development.
# Stripe does require webhooks, as we rely on them in order to know when a subscription becomes active.
if [[ ! -z "${SIMFINY_STRIPE_API_KEY}" ]] && \
  [[ ! -z "${SIMFINY_STRIPE_WEBHOOK_SECRET}" ]] && \
  [[ ! -z "${SIMFINY_STRIPE_DEFAULT_PRICE_ID}" ]] && \
  [[ ! -z "${WEBHOOKS_DOMAIN}" ]]; then
  echo "[wrapper] Stripe credentials are available, stripe and billing will be enabled...";
  export SIMFINY_STRIPE_ENABLED="true";
  export SIMFINY_STRIPE_BILLING_ENABLED="true";
  export SIMFINY_STRIPE_WEBHOOKS_ENABLED="true";
  export SIMFINY_STRIPE_WEBHOOKS_DOMAIN=${WEBHOOKS_DOMAIN};
fi

# Sometimes the old process does not get killed properly. This should do it.
pkill monetr;
pkill dlv;

# Execute the command with the new environment variables.
/go/bin/dlv exec --continue --api-version 2 --accept-multiclient --listen=:2345 --headless=true --api-version=2 /usr/bin/simfiny;
