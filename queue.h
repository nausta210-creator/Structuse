#ifndef QUEUE_H
#define QUEUE_H

#include <string>

using namespace std;

struct QNode {
    string data;
    QNode* next;
};

struct Queue {
    QNode* head;
    QNode* tail;
};

void initQueue(Queue& q);
void destroyQueue(Queue& q);
void pushQueue(Queue& q, const string& val);
string popQueue(Queue& q);
void printQueue(const Queue& q);

#endif