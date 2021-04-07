from django.core.serializers import serialize
from django.http import HttpResponse
from clients.models import Client

def index(request):
    clients = Client.objects.all()
    data = serialize("json", clients)
    return HttpResponse(data, content_type="application/json")
