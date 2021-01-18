# The values in this file will be overwritten on the airflow server with
# values pertaining to the image that we're build during the CI process
# and the branch from which they were built.

# Do not place any more configuration in this file - it will be removed
# on the server

DAG_NAME = 'some_dag'
IMAGES = {
        "task_1": "874568069660.dkr.ecr.eu-west-1.amazonaws.com/dst/some_dag_task1:latest",
        "task_2": "874568069660.dkr.ecr.eu-west-1.amazonaws.com/dst/some_dag_task2:latest"
}
