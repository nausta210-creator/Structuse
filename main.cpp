#include <iostream>
#include <string>
#include <sstream>
#include <vector>
#include <cctype>
#include "storage.h"
#include "array.h"
#include "flist.h"
#include "dlist.h"
#include "stack.h"
#include "queue.h"
#include "rbtree.h"

using namespace std;

vector<string> splitQuery(const string& query) {
    vector<string> tokens;
    stringstream ss(query);
    string token;
    while (ss >> token) tokens.push_back(token);
    return tokens;
}

bool isNumber(const string& s) {
    if (s.empty()) return false;
    for (size_t i = 0; i < s.size(); ++i) {
        if (!isdigit(s[i])) return false;
    }
    return true;
}

int main(int argc, char* argv[]) {
    string filename = "", query = "";

    for (int i = 1; i < argc; ++i) {
        string arg = argv[i];
        if (arg == "--file" && i + 1 < argc) filename = argv[++i];
        else if (arg == "--query" && i + 1 < argc) query = argv[++i];
    }

    if (filename.empty() || query.empty()) {
        cerr << "Использование: ./dbms --file <file.data> --query '<COMMAND> <name> [args]'" << endl;
        return 1;
    }

    vector<DataRecord> records = loadFromFile(filename);
    vector<string> args = splitQuery(query);
    if (args.empty()) return 0;

    string command = args[0];
    
    string structName = "";
    if (args.size() > 1) {
        structName = args[1];
    }

    string structType = "";
    if (command[0] == 'M') structType = "M";
    else if (command[0] == 'F') structType = "F";
    else if (command[0] == 'L') structType = "L";
    else if (command[0] == 'S') structType = "S";
    else if (command[0] == 'Q') structType = "Q";
    else if (command[0] == 'T') structType = "T";
    else if (command == "PRINT") {
        for (const auto& rec : records) {
            if (rec.name == structName) {
                structType = rec.type;
                break;
            }
        }
    }

    if (structType.empty()) {
        cerr << "Неизвестная команда или структура: " << command << endl;
        return 1;
    }

    bool modified = false;

    //МАССИВ
    if (structType == "M") {
        int recIdx = -1;
        for (size_t i = 0; i < records.size(); ++i) {
            if (records[i].name == structName && records[i].type == "M") { 
                recIdx = static_cast<int>(i); 
                break; 
            }
        }
        DynamicArray arr; initArray(arr);
        if (recIdx != -1) { 
            for (const auto& val : records[recIdx].items) {
                pushBack(arr, val); 
            }
        }

        if (command == "MPUSH") {
            if (args.size() == 3) { 
                pushBack(arr, args[2]); 
                modified = true; 
            }
            else if (args.size() == 4) { 
                insertAt(arr, stoi(args[2]), args[3]); 
                modified = true; 
            }
        } else if (command == "MDEL" && args.size() == 3) {
            removeAt(arr, stoi(args[2])); 
            modified = true;
        } else if (command == "MGET" && args.size() == 3) {
            cout << getAt(arr, stoi(args[2])) << endl;
        } else if (command == "MSET" && args.size() == 4) {
            setAt(arr, stoi(args[2]), args[3]); 
            modified = true;
        } else if (command == "MLEN") {
            cout << getSize(arr) << endl;
        } else if (command == "PRINT") {
            printArray(arr);
        }

        if (modified) {
            vector<string> newItems;
            for (size_t i = 0; i < arr.length; ++i) {
                newItems.push_back(arr.data[i]);
            }
            if (recIdx != -1) {
                records[recIdx].items = newItems;
            } else {
                records.push_back({"M", structName, newItems});
            }
            saveToFile(filename, records);
        }
        destroyArray(arr);
    }
    //ОДНОСВЯЗНЫЙ СПИСОК
    else if (structType == "F") {
        int recIdx = -1;
        for (size_t i = 0; i < records.size(); ++i) {
            if (records[i].name == structName && records[i].type == "F") { 
                recIdx = static_cast<int>(i); 
                break; 
            }
        }
        ForwardList list; initFList(list);
        if (recIdx != -1) { 
            for (const auto& val : records[recIdx].items) {
                pushTail(list, val); 
            }
        }

        if (command == "FPUSH" && args.size() >= 4) {
            string mode = args[2];
            if (mode == "HEAD") pushHead(list, args[3]);
            else if (mode == "TAIL") pushTail(list, args[3]);
            else if (mode == "AFTER" && args.size() == 5) pushAfter(list, args[3], args[4]);
            else if (mode == "BEFORE" && args.size() == 5) pushBefore(list, args[3], args[4]);
            modified = true;
        } else if (command == "FDEL" && args.size() >= 3) {
            string mode = args[2];
            if (mode == "HEAD") popHead(list);
            else if (mode == "TAIL") popTail(list);
            else if (mode == "VAL" && args.size() == 4) removeValue(list, args[3]);
            else if (mode == "BEFORE" && args.size() == 4) removeBefore(list, args[3]);
            else if (mode == "AFTER" && args.size() == 4) removeAfter(list, args[3]);
            modified = true;
        } else if (command == "FGET" && args.size() == 3) {
            if (isNumber(args[2])) {
                string val = getAtFList(list, stoi(args[2]));
                if (!val.empty()) {
                    cout << val << endl;
                }
            } else {
                bool found = findValue(list, args[2]);
                if (found) {
                    cout << "TRUE" << endl;
                } else {
                    cout << "FALSE" << endl;
                }
            }
        } else if (command == "PRINT") {
            printFList(list);
        } else if (command == "FPRINT_REC") {
            printFListRecursive(list);
        }

        if (modified) {
            vector<string> newItems;
            FNode* curr = list.head;
            while (curr != nullptr) { 
                newItems.push_back(curr->data); 
                curr = curr->next; 
            }
            if (recIdx != -1) {
                records[recIdx].items = newItems;
            } else {
                records.push_back({"F", structName, newItems});
            }
            saveToFile(filename, records);
        }
        destroyFList(list);
    }
    //ДВУСВЯЗНЫЙ СПИСОК
    else if (structType == "L") {
        int recIdx = -1;
        for (size_t i = 0; i < records.size(); ++i) {
            if (records[i].name == structName && records[i].type == "L") { 
                recIdx = static_cast<int>(i); 
                break; 
            }
        }
        DoublyLinkedList list; initDList(list);
        if (recIdx != -1) { 
            for (const auto& val : records[recIdx].items) {
                pushTailDList(list, val); 
            }
        }

        if (command == "LPUSH" && args.size() >= 4) {
            string mode = args[2];
            if (mode == "HEAD") pushHeadDList(list, args[3]);
            else if (mode == "TAIL") pushTailDList(list, args[3]);
            else if (mode == "AFTER" && args.size() == 5) pushAfterDList(list, args[3], args[4]);
            else if (mode == "BEFORE" && args.size() == 5) pushBeforeDList(list, args[3], args[4]);
            modified = true;
        } else if (command == "LDEL" && args.size() >= 3) {
            string mode = args[2];
            if (mode == "HEAD") popHeadDList(list);
            else if (mode == "TAIL") popTailDList(list);
            else if (mode == "VAL" && args.size() == 4) removeValueDList(list, args[3]);
            else if (mode == "BEFORE" && args.size() == 4) removeBeforeDList(list, args[3]);
            else if (mode == "AFTER" && args.size() == 4) removeAfterDList(list, args[3]);
            modified = true;
        } else if (command == "LGET" && args.size() == 3) {
            if (isNumber(args[2])) {
                string val = getAtDList(list, stoi(args[2]));
                if (!val.empty()) {
                    cout << val << endl;
                }
            } else {
                bool found = findValueDList(list, args[2]);
                if (found) {
                    cout << "TRUE" << endl;
                } else {
                    cout << "FALSE" << endl;
                }
            }
        } else if (command == "PRINT") {
            printDList(list);
        } else if (command == "LPRINT_REV") {
            printDListReverse(list);
        }

        if (modified) {
            vector<string> newItems;
            DNode* curr = list.head;
            while (curr != nullptr) { 
                newItems.push_back(curr->data); 
                curr = curr->next; 
            }
            if (recIdx != -1) {
                records[recIdx].items = newItems;
            } else {
                records.push_back({"L", structName, newItems});
            }
            saveToFile(filename, records);
        }
        destroyDList(list);
    }
    //СТЕК
    else if (structType == "S") {
        int recIdx = -1;
        for (size_t i = 0; i < records.size(); ++i) {
            if (records[i].name == structName && records[i].type == "S") { 
                recIdx = static_cast<int>(i); 
                break; 
            }
        }
        Stack st; initStack(st);
        if (recIdx != -1) {
            for (auto it = records[recIdx].items.rbegin(); it != records[recIdx].items.rend(); ++it) {
                pushStack(st, *it);
            }
        }

        if (command == "SPUSH" && args.size() == 3) {
            pushStack(st, args[2]);
            modified = true;
        } else if (command == "SPOP") {
            string val = popStack(st);
            if (!val.empty()) {
                cout << val << endl;
            }
            modified = true;
        } else if (command == "PRINT") {
            printStack(st);
        }

        if (modified) {
            vector<string> newItems;
            SNode* curr = st.top;
            while (curr != nullptr) { 
                newItems.push_back(curr->data); 
                curr = curr->next; 
            }
            if (recIdx != -1) {
                records[recIdx].items = newItems;
            } else {
                records.push_back({"S", structName, newItems});
            }
            saveToFile(filename, records);
        }
        destroyStack(st);
    }
    //ОЧЕРЕДЬ
    else if (structType == "Q") {
        int recIdx = -1;
        for (size_t i = 0; i < records.size(); ++i) {
            if (records[i].name == structName && records[i].type == "Q") { 
                recIdx = static_cast<int>(i); 
                break; 
            }
        }
        Queue q; initQueue(q);
        if (recIdx != -1) {
            for (const auto& val : records[recIdx].items) {
                pushQueue(q, val);
            }
        }

        if (command == "QPUSH" && args.size() == 3) {
            pushQueue(q, args[2]);
            modified = true;
        } else if (command == "QPOP") {
            string val = popQueue(q);
            if (!val.empty()) {
                cout << val << endl;
            }
            modified = true;
        } else if (command == "PRINT") {
            printQueue(q);
        }

        if (modified) {
            vector<string> newItems;
            QNode* curr = q.head;
            while (curr != nullptr) { 
                newItems.push_back(curr->data); 
                curr = curr->next; 
            }
            if (recIdx != -1) {
                records[recIdx].items = newItems;
            } else {
                records.push_back({"Q", structName, newItems});
            }
            saveToFile(filename, records);
        }
        destroyQueue(q);
    }
    //КРАСНО-ЧЕРНОЕ ДЕРЕВО
    else if (structType == "T") {
        int recIdx = -1;
        for (size_t i = 0; i < records.size(); ++i) {
            if (records[i].name == structName && records[i].type == "T") { 
                recIdx = static_cast<int>(i); 
                break; 
            }
        }
        
        RBTree tree; 
        initTree(tree);
        
        if (recIdx != -1) {
            for (const auto& val : records[recIdx].items) {
                insertNode(tree, stoi(val));      
            }
        }

        if (command == "TINSERT" && args.size() == 3) {
            insertNode(tree, stoi(args[2]));       
            modified = true;
        } else if (command == "TDEL" && args.size() == 3) {
            deleteNode(tree, stoi(args[2]));      
            modified = true;
        } else if (command == "TGET" && args.size() == 3) {
            bool found = searchNode(tree, stoi(args[2]));  
            if (found) {
                cout << "TRUE" << endl;
            } else {
                cout << "FALSE" << endl;
            }
        } else if (command == "PRINT") {
            printTree(tree);
        }

        vector<string> newItems;
        extractElements(tree, newItems);

        bool treeChanged = false;
        if (recIdx == -1) {
            treeChanged = true;                 
        } else if (records[recIdx].items != newItems) {
            treeChanged = true;                 
        }

        if (modified || treeChanged) {
            if (recIdx != -1) {
                records[recIdx].items = newItems;
            } else {
                records.push_back({"T", structName, newItems});
            }
            saveToFile(filename, records);
        }
        
        destroyTree(tree);
    }

    return 0;
}