# Commands

## Reference

    whom init <repository name>
        Initialize a new repository at ./<repository name>/

    whom pod new <pod name>
        Creates a new pod to manage.

        Options:
            --command,-c
                Use a podman command to run the pod.

            --compose,-m
                Use a compose.yml file to run the pod.

            --edit,-e
                If success, then open the pod file with $EDITOR

    whom pod rm <pod name>
        Removes a pod from the repository.

        Options:
            --danger-rm-all
                Delete all the pods in the current repository. Permanent action.

    whom pod edit <pod name>
        Open the pod file to edit it using $EDITOR.

    whom pod start <pod name>
        Activate the pod systemd service.

        Options:
            --all
                Start all pods in the repository.

    whom pod stop <pod name>
        Deactivate the pod systemd service.

        Options:
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

