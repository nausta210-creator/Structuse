#ifndef RBTREE_H
#define RBTREE_H

#include <string>
#include <vector>

using namespace std;

enum RBTColor { BLACK, RED };

struct RBTNode {
    int data;
    RBTColor color;
    RBTNode* left;
    RBTNode* right;
    RBTNode* parent;
};

struct RBTree {
    RBTNode* root;
    RBTNode* nil;
};

void initTree(RBTree& t);
void destroyTree(RBTree& t);

void insertNode(RBTree& t, int  key);
void deleteNode(RBTree& t, int  key);
bool searchNode(const RBTree& t, int key);

void printTree(const RBTree& t);
void extractElements(const RBTree& t, vector<string>& outItems);

#endif