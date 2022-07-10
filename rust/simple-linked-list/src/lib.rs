use std::{iter::FromIterator};

pub struct SimpleLinkedList<T> {
    head: Option<Box<Node<T>>>,
}

struct Node<T> {
    data: T,
    next: Option<Box<Node<T>>>,
}

impl<T> SimpleLinkedList<T> {
    pub fn new() -> Self {
        SimpleLinkedList { head: None }
    }

    pub fn is_empty(&self) -> bool {
        self.head.is_none()
    }

    pub fn len(&self) -> usize {
        let mut len: usize = 0;
        let mut current_node = &self.head;
        while let Some(node) = current_node {
            len += 1;
            current_node = &node.next;
        }
        len
    }

    pub fn push(&mut self, element: T) {
        let new_node = Some(Box::new(Node {
            data: element,
            next: self.head.take(),
        }));
        self.head = new_node;
    }

    pub fn pop(&mut self) -> Option<T> {
        match self.head.take() {
            None => None,
            Some(node) => {
                self.head = node.next;
                Some(node.data)
            }
        }
    }

    pub fn peek(&self) -> Option<&T> {
        match self.head {
            None => None,
            Some(ref current_node) => Some(&current_node.data),
        }
    }

    #[must_use]
    pub fn rev(mut self) -> SimpleLinkedList<T> {
        let mut new_list = SimpleLinkedList::new();
        while let Some(element) = self.pop() {
            new_list.push(element)
        }
        new_list
    }
}

impl<T> FromIterator<T> for SimpleLinkedList<T> {
    fn from_iter<I: IntoIterator<Item = T>>(iter: I) -> Self {
        let mut new_list = SimpleLinkedList::new();
        for value in iter {
            new_list.push(value)
        }
        new_list
    }
}

impl<T> From<SimpleLinkedList<T>> for Vec<T> {
    fn from(mut linked_list: SimpleLinkedList<T>) -> Vec<T> {
        let mut vector: Vec<T> = Vec::new();
        while let Some(element) = linked_list.pop() {
            vector.insert(0, element);
        }
        vector
    }
}
