#include "storage.h"
#include <fstream>
#include <sstream>
#include <iostream>

using namespace std;

vector<DataRecord> loadFromFile(const string& filename) {
    vector<DataRecord> records;
    ifstream file(filename);

    if (!file.is_open()) {
        return records;
    }

    string line;
    while (getline(file, line)) {
        if (line.empty()) continue;

        stringstream ss(line);
        DataRecord record;
        ss >> record.type >> record.name;

        string item;
        while (ss >> item) {
            record.items.push_back(item);
        }

        records.push_back(record);
    }

    file.close();
    return records;
}

void saveToFile(const string& filename, const vector<DataRecord>& records) {
    ofstream file(filename);
    if (!file.is_open()) {
        cerr << "Ошибка открытия файла для записи: " << filename << endl;
        return;
    }

    for (const auto& record : records) {
        file << record.type << " " << record.name;
        for (const auto& item : record.items) {
            file << " " << item;
        }
        file << "\n";
    }

    file.close();
}