import boto3
import json
import os
from collections import namedtuple


def generate_message(parsed_record):
    return json.dumps(parsed_record._asdict())


def send_message(sqs_client, sqs_queue_url, answer_id, message):
    response = sqs_client.send_message(
        QueueUrl=sqs_queue_url,
        DelaySeconds=10,
        MessageAttributes={
            'Title': {
                'DataType': 'String',
                'StringValue': answer_id
            }
        },
        MessageBody=(
            message
        )
    )
    return response


def main():
    filepath = os.environ.get('DATA_FILE')
    ParsedRow = namedtuple('ParsedRow', 'case_id sex age country answers')
    positions = {
        'case_id': [0, 6],
        'sex': [6, 7],
        'age': [7, 9],
        'country': [22, 32],
        'answers': [33, 333]}

    ps_records = []

    with open(filepath) as fp:
        for cnt, line in enumerate(fp):
            parsed = ParsedRow(
                case_id=line[0:6].strip(),
                sex=line[6:7].strip(),
                age=line[7:9].strip(),
                country=line[22:32].strip(),
                answers=line[33:333].strip()
            )
            ps_records.append(parsed)

    aws_url = "http://localhost:4566"
    queue_name = 'results-queue'
    queue_url = "{}/000000000000/{}".format(aws_url, queue_name)
    aws_client = boto3.client('sqs', endpoint_url=aws_url)

    for record in ps_records:
        answer_id = record.case_id
        msg = generate_message(record)
        # ENV VAR
        response = send_message(aws_client, queue_url, answer_id, msg)
        print("sent message: msg_id: {} answer_id: {}".format(response['MessageId'], answer_id))


if __name__ == "__main__":
    main()
