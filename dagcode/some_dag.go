from datetime import timedelta
from datetime import datetime
from airflow import DAG
from airflow.utils.dates import days_ago
from etltoolkit.generators.kubernetes_generator import KubernetesGenerator
from etltoolkit.helpers.kube import res
from etltoolkit.macros import CONN_MACROS

import dags.dagcode.conf.conf as conf
import dags.dagcode.conf.cdconfig as cdconfig
from params import vars as v
from params import kube as k

# define DAG
dag = DAG(
    cdconfig.DAG_NAME,
    start_date=datetime(2020, 10, 14),
    schedule_interval="5 15 * * *",
    user_defined_macros=CONN_MACROS,
)

# Kubernetes generator
kube_gen = KubernetesGenerator(
        dag=dag,
        alert_level="non-critical",
)

task_1 = kube_gen.run_pod(
    name="task_1",
    image=cdconfig.IMAGES['task_1'],
    resources=res(0.01, 0.1),
    env_vars=conf.TASK1_ENVS,
    execution_timeout=timedelta(minutes=10),
)

task_2 = kube_gen.run_pod(
    name="task_2",
    image=cdconfig.IMAGES['task_2'],
    resources=res(0.01, 0.1),
    env_vars=conf.TASK2_ENVS,
    execution_timeout=timedelta(minutes=10),
)
