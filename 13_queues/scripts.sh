# redis

siege -d0.5 -v -t2s -250 --content-type 'application/json' 'http://localhost:8989/redis/test POST {"message": "some msg here"}'

siege -d0.5 -v -t2s -c250 'http://localhost:8989/redis/test'

# beanstalkd

siege -d0.5 -v -t2s -250 --content-type 'application/json' 'http://localhost:8989/beanstalk/test POST {"message": "some msg here"}'

siege -d0.5 -v -t2s -c250 'http://localhost:8989/beanstalk/test'
