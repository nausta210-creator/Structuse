#include "rbtree.h"
#include <iostream>

using namespace std;

void initTree(RBTree& t) {
    t.nil = new RBTNode{0, BLACK, nullptr, nullptr, nullptr};
    t.root = t.nil;
}

void destroyHelper(RBTNode* node, RBTNode* nil) {
    if (node != nil) {
        destroyHelper(node->left, nil);
        destroyHelper(node->right, nil);
        delete node;
    }
}

void destroyTree(RBTree& t) {
    destroyHelper(t.root, t.nil);
    delete t.nil;
}

void rotateLeft(RBTree& t, RBTNode* x) {
    RBTNode* y = x->right;
    x->right = y->left;
    if (y->left != t.nil) y->left->parent = x;
    y->parent = x->parent;
    if (x->parent == t.nil) t.root = y;
    else if (x == x->parent->left) x->parent->left = y;
    else x->parent->right = y;
    y->left = x;
    x->parent = y;
}

void rotateRight(RBTree& t, RBTNode* x) {
    RBTNode* y = x->left;
    x->left = y->right;
    if (y->right != t.nil) y->right->parent = x;
    y->parent = x->parent;
    if (x->parent == t.nil) t.root = y;
    else if (x == x->parent->right) x->parent->right = y;
    else x->parent->left = y;
    y->right = x;
    x->parent = y;
}

void fixInsert(RBTree& t, RBTNode* k) {
    while (k->parent->color == RED) {
        if (k->parent == k->parent->parent->left) {
            RBTNode* u = k->parent->parent->right;
            if (u->color == RED) {
                k->parent->color = BLACK;
                u->color = BLACK;
                k->parent->parent->color = RED;
                k = k->parent->parent;
            } else {
                if (k == k->parent->right) {
                    k = k->parent;
                    rotateLeft(t, k);
                }
                k->parent->color = BLACK;
                k->parent->parent->color = RED;
                rotateRight(t, k->parent->parent);
            }
        } else {
            RBTNode* u = k->parent->parent->left;
            if (u->color == RED) {
                k->parent->color = BLACK;
                u->color = BLACK;
                k->parent->parent->color = RED;
                k = k->parent->parent;
            } else {
                if (k == k->parent->left) {
                    k = k->parent;
                    rotateRight(t, k);
                }
                k->parent->color = BLACK;
                k->parent->parent->color = RED;
                rotateLeft(t, k->parent->parent);
            }
        }
    }
    t.root->color = BLACK;
}

void insertNode(RBTree& t, int key) {
    RBTNode* node = new RBTNode{key, RED, t.nil, t.nil, t.nil};
    RBTNode* y = t.nil;
    RBTNode* x = t.root;

    while (x != t.nil) {
        y = x;
        if (node->data == x->data) {
            delete node;
            return;
        }
        if (node->data < x->data) x = x->left;
        else x = x->right;
    }

    node->parent = y;
    if (y == t.nil) t.root = node;
    else if (node->data < y->data) y->left = node;
    else y->right = node;

    if (node->parent == t.nil) {
        node->color = BLACK;
        return;
    }
    if (node->parent->parent == t.nil) return;

    fixInsert(t, node);
}

void transplant(RBTree& t, RBTNode* u, RBTNode* v) {
    if (u->parent == t.nil) t.root = v;
    else if (u == u->parent->left) u->parent->left = v;
    else u->parent->right = v;
    v->parent = u->parent;
}

RBTNode* minimumNode(RBTree& t, RBTNode* node) {
    while (node->left != t.nil) node = node->left;
    return node;
}

void fixDelete(RBTree& t, RBTNode* x) {
    while (x != t.root && x->color == BLACK) {
        if (x == x->parent->left) {
            RBTNode* s = x->parent->right;
            if (s->color == RED) {
                s->color = BLACK;
                x->parent->color = RED;
                rotateLeft(t, x->parent);
                s = x->parent->right;
            }
            if (s->left->color == BLACK && s->right->color == BLACK) {
                s->color = RED;
                x = x->parent;
            } else {
                if (s->right->color == BLACK) {
                    s->left->color = BLACK;
                    s->color = RED;
                    rotateRight(t, s);
                    s = x->parent->right;
                }
                s->color = x->parent->color;
                x->parent->color = BLACK;
                s->right->color = BLACK;
                rotateLeft(t, x->parent);
                x = t.root;
            }
        } else {
            RBTNode* s = x->parent->left;
            if (s->color == RED) {
                s->color = BLACK;
                x->parent->color = RED;
                rotateRight(t, x->parent);
                s = x->parent->left;
            }
            if (s->right->color == BLACK && s->left->color == BLACK) {
                s->color = RED;
                x = x->parent;
            } else {
                if (s->left->color == BLACK) {
                    s->right->color = BLACK;
                    s->color = RED;
                    rotateLeft(t, s);
                    s = x->parent->left;
                }
                s->color = x->parent->color;
                x->parent->color = BLACK;
                s->left->color = BLACK;
                rotateRight(t, x->parent);
                x = t.root;
            }
        }
    }
    x->color = BLACK;
}

void deleteNode(RBTree& t, int key) {
    RBTNode* z = t.root;
    while (z != t.nil && z->data != key) {
        if (key < z->data) z = z->left;
        else z = z->right;
    }
    if (z == t.nil) return;

    RBTNode* y = z;
    RBTNode* x;
    RBTColor yOriginalColor = y->color;

    if (z->left == t.nil) {
        x = z->right;
        transplant(t, z, z->right);
    } else if (z->right == t.nil) {
        x = z->left;
        transplant(t, z, z->left);
    } else {
        y = minimumNode(t, z->right);
        yOriginalColor = y->color;
        x = y->right;
        if (y->parent == z) {
            x->parent = y;
        } else {
            transplant(t, y, y->right);
            y->right = z->right;
            y->right->parent = y;
        }
        transplant(t, z, y);
        y->left = z->left;
        y->left->parent = y;
        y->color = z->color;
    }
    delete z;
    if (yOriginalColor == BLACK) fixDelete(t, x);
}

bool searchNode(const RBTree& t, int key) {
    RBTNode* curr = t.root;
    while (curr != t.nil) {
        if (key == curr->data) return true;
        if (key < curr->data) curr = curr->left;
        else curr = curr->right;
    }
    return false;
}

void printHelper(RBTNode* root, RBTNode* nil, string indent, bool last) {
    if (root != nil) {
        cout << indent;
        if (last) {
            cout << "R----";
            indent += "     ";
        } else {
            cout << "L----";
            indent += "|    ";
        }
        
        string sColor = "(BLK)";
        if (root->color == RED) {
            sColor = "(RED)";
        }
        
        cout << root->data << " " << sColor << endl;
        printHelper(root->left, nil, indent, false);
        printHelper(root->right, nil, indent, true);
    }
}

void printTree(const RBTree& t) {
    if (t.root == t.nil) {
        cout << "[Пустое дерево]" << endl;
        return;
    }
    printHelper(t.root, t.nil, "", true);
}

void extractHelper(RBTNode* node, RBTNode* nil, vector<string>& outItems) {
    if (node != nil) {
        outItems.push_back(to_string(node->data));
        extractHelper(node->left, nil, outItems);
        extractHelper(node->right, nil, outItems);
    }
}

void extractElements(const RBTree& t, vector<string>& outItems) {
    extractHelper(t.root, t.nil, outItems);
}