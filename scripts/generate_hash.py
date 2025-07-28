import hashlib

password = "Hello_World!"
hash_object = hashlib.sha512(password.encode())
print(hash_object.hexdigest())