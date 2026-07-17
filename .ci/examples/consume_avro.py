import pulsar
from pulsar.schema import *

from _pulsar import InitialPosition

import sys

topic = sys.argv[1]
expected_name = sys.argv[2] if len(sys.argv) > 2 else None
expected_age = int(sys.argv[3]) if len(sys.argv) > 3 else None
expected_grade = int(sys.argv[4]) if len(sys.argv) > 4 else None


class Student(Record):
    def __init__(self, name, age, grade):
        self.name = name
        self.age = age
        self.grade = grade

    name = String()
    age = Integer()
    grade = Integer()


client = pulsar.Client('pulsar://localhost:6650')

schema = pulsar.schema.AvroSchema(Student)
consumer = client.subscribe(topic=topic,
                            subscription_name='my-avro-subscription',
                            crypto_key_reader=None,
                            initial_position=InitialPosition.Earliest,
                            schema=schema)

try:
    msg = consumer.receive(1000)
    value = msg.value()
    if expected_name is not None:
        if value.name != expected_name or value.age != expected_age or value.grade != expected_grade:
            print(value)
            sys.exit(1)
        print("schema message matches")
    else:
        print(value)
except Exception:
    if expected_name is not None:
        sys.exit(1)
