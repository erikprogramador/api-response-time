from django.db import models

class Client(models.Model):
    name = models.CharField(max_length=150)
    email = models.CharField(max_length=150)
    phone = models.CharField(max_length=50)
    domain = models.CharField(max_length=255)
    description = models.TextField()
    tags = models.CharField(max_length=255)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    class Meta:
        db_table = "clients"

    def __str__(self):
        return self.name
