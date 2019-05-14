import json
from aliyunsdkcore.client import AcsClient
import sys

from aliyunsdkslb.request.v20140515 import SetBackendServersRequest


from aliyunsdkcore import client
from aliyunsdkslb.request.v20140515 import DescribeLoadBalancerAttributeRequest
from aliyunsdkcore.profile import region_provider


clt = client.AcsClient("LTAIqM1yei0xXdQ7","0T01csJF2SvRsFybw9X1g8s6dGxDO9","cn-shenzhen")

request = DescribeLoadBalancerAttributeRequest.DescribeLoadBalancerAttributeRequest()
request.set_accept_format('json')

request.add_query_param('RegionId', 'cn-shenzhen')
request.add_query_param('LoadBalancerId', 'lb-wz9bic77x0m0xfdtnjy0t')

response = clt.do_action_with_exception(request)
weight = json.loads(response)["BackendServers"]["BackendServer"]
weight[0]["Weight"] = 100
weight[1]["Weight"] = 100
# for i in range(len(weight)):
#     weight[i]["Weight"] = sys.argv[i+1]
#     print(weight[i]["Weight"])


weight = json.dumps(weight)


request = SetBackendServersRequest.SetBackendServersRequest()
request.set_accept_format('json')


request.add_query_param('RegionId', 'cn-shenzhen')
request.add_query_param('LoadBalancerId', 'lb-wz9bic77x0m0xfdtnjy0t')
request.add_query_param('BackendServers', weight)



response = clt.do_action(request)




print(response,'\n')