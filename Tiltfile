# Welcome to Tilt!
#   To get you started as quickly as possible, we have created a
#   starter Tiltfile for you.
#
#   Uncomment, modify, and delete any commands as needed for your
#   project's configuration.

# Extensions are open-source, pre-packaged functions that extend Tilt
#
#   More info: https://github.com/tilt-dev/tilt-extensions
#
load('ext://git_resource', 'git_checkout')

## includes cancel button for any resource in the UI
include('ext://cancel')

## deploy cert-manager to cluster
load('ext://cert_manager', 'deploy_cert_manager')
deploy_cert_manager()

## add ANSI color codes to tiltfile log messages
load('ext://color', 'color')

## Use this extension when you have an image and you want to re-execute its entrypoint/command as part of a live_update.
load('ext://restart_process', 'docker_build_with_restart')

## kubectl_build integrates BuildKit CLI for kubectl, a tool for building OCI and Docker images with your kubernetes cluster, with tilt.
load('ext://kubectl_build', 'kubectl_build')

# When you include this extension, Tilt starts an ngrok server on port 4040 with no tunnels open by default.
#
# Every service with a link or port-forward will have a new button labelled 'ngrok'.
#
# Clicking the button will create a new tunnel. You can see all the currently open tunnels at http://localhost:4040/.
# ref: https://github.com/tilt-dev/tilt-extensions/tree/master/ngrok
v1alpha1.extension_repo(name='default', url='https://github.com/tilt-dev/tilt-extensions')
v1alpha1.extension(name='ngrok:config', repo_name='default', repo_path='ngrok')

# Output diagnostic messages
#   You can print log messages, warnings, and fatal errors, which will
#   appear in the (Tiltfile) resource in the web UI. Tiltfiles support
#   multiline strings and common string operations such as formatting.
#
#   More info: https://docs.tilt.dev/api.html#api.warn
print("""
-----------------------------------------------------------------
✨ Hello Tilt! This appears in the (Tiltfile) pane whenever Tilt
   evaluates this file.
-----------------------------------------------------------------
""".strip())
warn('ℹ️ Open {tiltfile_path} in your favorite editor to get started.'.format(
    tiltfile_path=config.main_path))


# force image rebuilds to happen with a parellelism level of 10
update_settings(max_parallel_updates=10)


# docker_build('registry.example.com/my-image',
#              context='.',
#              # (Optional) Use a custom Dockerfile path
#              dockerfile='./deploy/app.dockerfile',
#              # (Optional) Filter the paths used in the build
#              only=['./app'],
#              # (Recommended) Updating a running container in-place
#              # https://docs.tilt.dev/live_update_reference.html
#              live_update=[
#                 # Sync files from host to container
#                 sync('./app', '/src/'),
#                 # Execute commands inside the container when certain
#                 # paths change
#                 run('/src/codegen.sh', trigger=['./app/api'])
#              ]
# )

# Run local commands
#   Local commands can be helpful for one-time tasks like installing
#   project prerequisites. They can also manage long-lived processes
#   for non-containerized services or dependencies.
#
#   More info: https://docs.tilt.dev/local_resource.html
local_resource('create-local-namespace',
                cmd="kubectl create ns local"
)

# Build Docker image
#   Tilt will automatically associate image builds with the resource(s)
#   that reference them (e.g. via Kubernetes or Docker Compose YAML).
#
#   More info: https://docs.tilt.dev/api.html#api.docker_build
#docker_build_with_restart
kubectl_build('feelguuds/financial-integration-service:latest',
              context='.',
              # (Optional) Use a custom Dockerfile path
              dockerfile='./Dockerfile.dev',
              # (Recommended) Updating a running container in-place
              # https://docs.tilt.dev/live_update_reference.html
              live_update=[
                 # Sync files from host to container
                 sync('./pkg', '/go/src/github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg'),
                 sync('./rpc', '/go/src/github.com/SimifiniiCTO/simfiny-financial-integration-service/proto'),
                 # Execute commands inside the container when certain
                 # paths change
                 #run('/src/codegen.sh', trigger=['./app/api'])
              ]
 )

# Apply Kubernetes manifests
#   Tilt will build & push any necessary images, re-deploying your
#   resources as they change.
#
#   More info: https://docs.tilt.dev/api.html#api.k8s_yaml
#
yaml = helm(
  './charts/financial-integration-service',
  # The release name, equivalent to helm --name
  name='service',
  # The namespace to install in, equivalent to helm --namespace
  namespace='local',
  # The values file to substitute into the chart.
  values=['./charts/financial-integration-service/values.yaml'],
  # Values to set from the command-line - we disable linkerd sidecar for local development for now
  set=['linkerd.profile.enabled=true']
  )
k8s_yaml(yaml)

k8s_resource(
    workload='service',
    port_forwards=[
      port_forward(9000, 9896, name='financial-integration-service-api'), 
   ]
    auto_init=True,
)


# Customize a Kubernetes resource
#   By default, Kubernetes resource names are automatically assigned
#   based on objects in the YAML manifests, e.g. Deployment name.
#
#   Tilt strives for sane defaults, so calling k8s_resource is
#   optional, and you only need to pass the arguments you want to
#   override.
#
#   More info: https://docs.tilt.dev/api.html#api.k8s_resource
#
# k8s_resource('my-deployment',
#              # map one or more local ports to ports on your Pod
#              port_forwards=['5000:8080'],
#              # change whether the resource is started by default
#              auto_init=False,
#              # control whether the resource automatically updates
#              trigger_mode=TRIGGER_MODE_MANUAL
# )


# Run local commands
#   Local commands can be helpful for one-time tasks like installing
#   project prerequisites. They can also manage long-lived processes
#   for non-containerized services or dependencies.
#
#   More info: https://docs.tilt.dev/local_resource.html
#
# local_resource('install-helm',
#                cmd='which helm > /dev/null || brew install helm',
#                # `cmd_bat`, when present, is used instead of `cmd` on Windows.
#                cmd_bat=[
#                    'powershell.exe',
#                    '-Noninteractive',
#                    '-Command',
#                    '& {if (!(Get-Command helm -ErrorAction SilentlyContinue)) {scoop install helm}}'
#                ]
# )


# Organize logic into functions
#   Tiltfiles are written in Starlark, a Python-inspired language, so
#   you can use functions, conditionals, loops, and more.
#
#   More info: https://docs.tilt.dev/tiltfile_concepts.html
#
def tilt_demo():
    # Tilt provides many useful portable built-ins
    # https://docs.tilt.dev/api.html#modules.os.path.exists
    if os.path.exists('tilt-avatars/Tiltfile'):
        # It's possible to load other Tiltfiles to further organize
        # your logic in large projects
        # https://docs.tilt.dev/multiple_repos.html
        load_dynamic('tilt-avatars/Tiltfile')
    watch_file('tilt-avatars/Tiltfile')
    git_checkout('https://github.com/tilt-dev/tilt-avatars.git',
                 checkout_dir='tilt-avatars')


# Edit your Tiltfile without restarting Tilt
#   While running `tilt up`, Tilt watches the Tiltfile on disk and
#   automatically re-evaluates it on change.
#
#   To see it in action, try uncommenting the following line with
#   Tilt running.
# tilt_demo()
