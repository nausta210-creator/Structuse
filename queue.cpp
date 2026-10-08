#include "queue.h"
#include <iostream>

using namespace std;

void initQueue(Queue& q) {
    q.head = nullptr;
    q.tail = nullptr;
}

void destroyQueue(Queue& q) {
    while (q.head != nullptr) {
        popQueue(q);
    }
}

void pushQueue(Queue& q, const string& val) {
    QNode* newNode = new QNode{val, nullptr};
    if (q.tail != nullptr) {
        q.tail->next = newNode;
    } else {
        q.head = newNode;
    }
    q.tail = newNode;
}

string popQueue(Queue& q) {
    if (q.head == nullptr) return "";
    QNode* temp = q.head;
    string val = temp->data;
    q.head = q.head->next;
    if (q.head == nullptr) {
        q.tail = nullptr;
    }
    delete temp;
    return val;
}

void printQueue(const Queue& q) {
    if (q.head == nullptr) {
        cout << "[Пустая очередь]" << endl;
        return;
    }
    cout << "[ ";
    QNode* curr = q.head;
    while (curr != nullptr) {
        cout << curr->data;
        if (curr->next != nullptr) {
            cout << " <- ";
        }
        curr = curr->next;
    }
    cout << " ]" << endl;
}