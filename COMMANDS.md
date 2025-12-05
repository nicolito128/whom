# Commands

## Reference

    whom init <repository name>
        Initialize a new repository at ./<repository name>/

    whom pod new
        Creates a new pod to manage.

        Options:
            --name,-n <pod name>
                Use a unique name for the new pod.

            --use,-u [ command, compose ]
                Select whether you want to use a plain text command file or a compose.yml file for the new pod.

    whom pod rm
        Removes a pod from the repository.

        Options:
            --name,-n <pod name>
                Refers to the pod to delete from the current repository.

            --danger-rm-all
                Delete all the pods in the current repository. Permanent action.

    whom pod start
        Enable the pod systemd service.

        Options:
            --name,-n <pod name>
                Start the <pod name> service.

            --all
                Start all pods in the repository.

    whom pod stop
        Disable the pod systemd service.

        Options:
            --name,-n <pod name>
                Stop the <pod name> service.

            --all
                Stop all pods in the repository.

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

