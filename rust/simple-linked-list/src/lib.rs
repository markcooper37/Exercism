use std::{iter::FromIterator, mem};

pub struct SimpleLinkedList<T> {
    head: Option<Box<Node<T>>>,
}

struct Node<T> {
    data: T,
    next: Option<Box<Node<T>>>,
}

impl<T: Copy> SimpleLinkedList<T> {
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

    pub fn push(&mut self, _element: T) {
        let new_node = Some(Box::new(Node {
            data: _element,
            next: mem::replace(&mut self.head, None),
        }));
        self.head = new_node;
    }

    pub fn pop(&mut self) -> Option<T> {
        match mem::replace(&mut self.head, None) {
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
    pub fn rev(self) -> SimpleLinkedList<T> {
        let mut new_list = SimpleLinkedList::new();
        let mut self_copy = self;
        while !self_copy.head.is_none() {
            new_list.push(self_copy.pop().unwrap())
        }
        new_list
    }
}

impl<T: Copy> FromIterator<T> for SimpleLinkedList<T> {
    fn from_iter<I: IntoIterator<Item = T>>(_iter: I) -> Self {
        let mut new_list = SimpleLinkedList::new();
        for value in _iter.into_iter() {
            new_list.push(value)
        }
        new_list
    }
}

impl<T: Copy> From<SimpleLinkedList<T>> for Vec<T> {
    fn from(mut _linked_list: SimpleLinkedList<T>) -> Vec<T> {
        let mut vector: Vec<T> = Vec::new();
        let mut current_node = &_linked_list.head;
        while let Some(node) = current_node {
            vector.insert(0, node.data);
            current_node = &node.next;
        }
        vector
    }
}
