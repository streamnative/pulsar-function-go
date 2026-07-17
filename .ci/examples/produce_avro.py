import pulsar
from pulsar.schema import *

import sys

topic = sys.argv[1]
name = sys.argv[2]
age = int(sys.argv[3])
grade = int(sys.argv[4])

client = pulsar.Client('pulsar://localhost:6650')


class Student(Record):
    def __init__(self, name, age, grade):
        self.name = name
        self.age = age
        self.grade = grade

    name = String()
    age = Integer()
    grade = Integer()


schema = pulsar.schema.AvroSchema(Student)
producer = client.create_producer(topic=topic,
                                  schema=schema)

student = Student(name, age, grade)
producer.send(student)

producer.close()
