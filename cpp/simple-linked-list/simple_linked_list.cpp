#include "simple_linked_list.h"

#include <stdexcept>
#include <iostream>

namespace simple_linked_list {

std::size_t List::size() const {
    return current_size;
}

void List::push(int entry) {
    auto new_element = new Element{entry};
    new_element->next = head;
    head = new_element;
    current_size++;
}

int List::pop() {
    if (head == nullptr) {
        throw std::runtime_error("List is empty");
    }
    int entry = head->data;
    head = head->next;
    current_size--;
    return entry;
}

void List::reverse() {
    Element* new_head = nullptr;
    while (head != nullptr) {
        auto temp = new_head;
        new_head = head;
        head = head->next;
        new_head->next = temp;
    }
    head = new_head;
}

List::~List() {
    while (head != nullptr) {
        Element* next = head->next;
        delete head;
        head = next;
    }
}

}  // namespace simple_linked_list
