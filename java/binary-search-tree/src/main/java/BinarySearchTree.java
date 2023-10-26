import java.util.List;
import java.util.ArrayList;

class BinarySearchTree<T extends Comparable<T>> {

    private Node<T> root;
    
    void insert(T value) {
        if (this.root == null) {
            this.root = new Node(value);
            return;
        };
        Node<T> parent = root;
        while (true) {
            int compRes = value.compareTo(parent.getData());
            if (compRes <= 0) {
                if (parent.getLeft() == null) {
                    parent.setLeft(value);
                    break;
                } else {
                    parent = parent.getLeft();
                    continue;
                }
            } else {
                if (parent.getRight() == null) {
                    parent.setRight(value);
                    break;
                } else {
                    parent = parent.getRight();
                    continue;
                }
            }
        }
    }

    List<T> getAsSortedList() {
        if (this.root == null) {
            return new ArrayList<T>();
        }
        return this.root.getAsSortedList();
    }

    List<T> getAsLevelOrderList() {
        List<T> values = new ArrayList<T>();
        List<Node<T>> levelNodes = new ArrayList<Node<T>>();
        if (this.root != null) {
            levelNodes.add(this.root);
        }
        while (levelNodes.size() != 0) {
            List<Node<T>> newLevel = new ArrayList<Node<T>>();
            levelNodes.forEach((node) -> {
                values.add(node.getData());
                if (node.getLeft() != null) {
                    newLevel.add(node.getLeft());
                }
                if (node.getRight() != null) {
                    newLevel.add(node.getRight());
                }
            });
            levelNodes = newLevel;
        }
        return values;
    }

    Node<T> getRoot() {
        return this.root;
    }

    static class Node<T> {

        private T data;
        private Node<T> left;
        private Node<T> right;

        public Node(T data) {
            this.data = data;
        }

        public List<T> getAsSortedList() {
            List<T> sortedList = this.left == null ? new ArrayList<T>() : this.left.getAsSortedList();
            sortedList.add(this.data);
            List<T> rightData = this.right == null ? new ArrayList<T>() : this.right.getAsSortedList();
            sortedList.addAll(rightData);
            return sortedList;
        }

        Node<T> getLeft() {
            return this.left;
        }

        public void setLeft(T data) {
            this.left = new Node(data);
        }

        Node<T> getRight() {
            return this.right;
        }

        public void setRight(T data) {
            this.right = new Node(data);
        }

        T getData() {
            return this.data;
        }

    }
}
