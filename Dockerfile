FROM ubuntu:14.04

ADD . /opt/app/

RUN apt-get update && apt-get install -y build-essential python-dev python-pip python-mysqldb

RUN pip install -r /opt/app/requirements.txt
RUN pip install MySQL-python

WORKDIR /opt/app
EXPOSE 80
CMD ["python", "run.py"]
