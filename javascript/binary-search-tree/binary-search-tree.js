export class BinarySearchTree {
  constructor(data) {
    this._data = data;
  }

  get data() {
    return this._data;
  }
  
  get right() {
    return this._right;
  }

  get left() {
    return this._left;
  }

  insert(data) {
    let current = this;
    while (true) {
      if (data <= current._data) {
        if (!current._left) {
          current._left = new BinarySearchTree(data);
          break;
        } else {
          current = current._left;
        };
      } else {
        if (!current._right) {
          current._right = new BinarySearchTree(data);
          break;
        } else {
          current = current._right;
        };
      };
    };
  }

  each(func) {
    if (this._left) {
      this._left.each(func);
    };
    func(this._data);
    if (this._right) {
      this._right.each(func);
    };
  }
}
