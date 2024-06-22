# Tweet Fanout Delivery Architecture Project

This project is built to understand Twitter's tweet fanout delivery architecture. The core idea is to simulate how tweets are delivered to followers using a message queue system. RabbitMQ is used as the message queue, and the Golang AMQP client is used to interact with RabbitMQ.

## Project Overview

The project provides the following endpoints:

1. **Create User**
2. **Follow User**
3. **Tweet**
4. **Get User Homepage**

### Architecture

1. **Create User**: Allows the creation of a new user in the system.
2. **Follow User**: Enables a user to follow an already existing user.
3. **Tweet**: When a user tweets, their tweet is placed in the message queues of all their followers.
4. **Get User Homepage**: Fetches all the tweets from the user's message queue, representing their homepage.

### Flow

- A user follows another user, subscribing to their tweets.
- When a user tweets, the tweet is published to the queues of all their followers.
- To retrieve the tweets, a user fetches messages from their own queue.
