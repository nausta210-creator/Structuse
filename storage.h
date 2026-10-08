#ifndef STORAGE_H
#define STORAGE_H

#include <string>
#include <vector>

using namespace std;

struct DataRecord {
    string type;          
    string name;          
    vector<string> items;
};

vector<DataRecord> loadFromFile(const string& filename);

void saveToFile(const string& filename, const vector<DataRecord>& records);

#endif