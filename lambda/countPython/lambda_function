from __future__ import print_function # Python 2/3 compatibility
import boto3
import json
import decimal
from boto3.dynamodb.conditions import Key, Attr
from botocore.exceptions import ClientError

class DecimalEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, decimal.Decimal):
            if o % 1 > 0:
                return float(o)
            else:
                return int(o)
        return super(DecimalEncoder, self).default(o)

dynamodb = boto3.resource("dynamodb", region_name='ap-northeast-2')

table = dynamodb.Table('visitors')

def lambda_handler(event, context):
    
    try:
        response = table.get_item(
            Key={
                'id': 1
            }
        )
    except ClientError as e:
        print(e.response['Error']['Message'])
        return {
            'statusCode': 404,
            'body': "NOT FOUND"
        }
    else:
        item = response['Item']
        print("GetItem succeeded:")
        print(json.dumps(item, indent=4, cls=DecimalEncoder))
        
        response = table.update_item(
            Key={
                'id': 1
            },
            ExpressionAttributeNames = {
                '#V': 'visitors'
            },
            ExpressionAttributeValues={
                ':i': 1
            },
            UpdateExpression="set #V = #V + :i",
            ReturnValues="UPDATED_NEW"
        )
        print("UpdateItem succeeded:")

        return {
            'statusCode': 200,
            'body': json.dumps(item, indent=4, cls=DecimalEncoder)
        }

