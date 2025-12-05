# Commands

## Reference

    whom init <repository name>
        Initialize a new repository at ./<repository name>/

    whom pod new
        Creates a new pod to manage.

        Options:
            --name,-n <pod name>
                Set a unique name for the new pod.

            --command,-c
                Use a podman command to run the pod.

            --compose,-f
                Use a compose.yml file to run the pod.

    whom pod rm
        Removes a pod from the repository.

        Options:
            --name,-n <pod name>
                Refers to the pod to delete from the current repository.

            --danger-rm-all
                Delete all the pods in the current repository. Permanent action.

    whom pod start
        Activate the pod systemd service.

        Options:
            --name,-n <pod name>
                Start the <pod name> service.

            --all
                Start all pods in the repository.

    whom pod stop
        Deactivate the pod systemd service.

        Options:
            --name,-n <pod name>
                Stop the <pod name> service.

            --all
                Stop all pods in the repository.

    whom pods
        List all active pods in the current repository.

        Options:
            --verbose,-v
                Show detailed information about each pod.

            --all
                Show all pods, including inactive ones.

## Project structure

    repository/
        pods/
            pod1/
                - command / compose.yml
            pod2/
                - command / compose.yml
            pod3/
            .
            .
            .
            podN/...

        config.toml
        pods.toml

