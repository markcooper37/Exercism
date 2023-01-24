export class LinkedList {
  push(value) {
    let newObject = {value: value};
    if (!this._first) {
      [this._first, this._last] = [newObject, newObject];
    } else {
      newObject.previous = this._last;
      this._last.next = newObject;
      this._last = newObject;
    };
  }

  pop() {
    let value = this._last.value;
    if (!this._last.previous) {
      delete this._first, this._last;
    } else {
      this._last = this._last.previous;
      delete this._last.next;
    };
    return value;
  }

  shift() {
    let value = this._first.value;
    if (!this._first.next) {
      delete this._first, this._last;
    } else {
      this._first = this._first.next;
      delete this._first.previous;
    };
    return value;
  }

  unshift(value) {
    let newObject = {value: value};
    if (!this._first) {
      [this._first, this._last] = [newObject, newObject];
    } else {
      newObject.next = this._first;
      this._first.previous = newObject;
      this._first = newObject;
    };
  }

  delete(value) {
    if (!this._first) {
      return;
    } else {
      let currentObject = this._first;
      while (true) {
        if (currentObject.value === value) {
          if (currentObject.previous && currentObject.next) {
            currentObject.previous.next = currentObject.next;
            currentObject.next.previous = currentObject.previous;
          } else if (currentObject.previous) {
            this._last = currentObject.previous;
            delete this._last.next;
          } else if (currentObject.next) {
            this._first = currentObject.next;
            delete this._first.previous;
          } else {
            delete this._first, this._last;
          };
          break
        } else if (!currentObject.next) {
          break;
        } else {
          currentObject = currentObject.next;
        };
      };
    };
  }

  count() {
    if (!this._first) {
      return 0;
    } else {
      let [count, currentObject] = [1, this._first];
      while (currentObject.next) {
        count++;
        currentObject = currentObject.next;
      };
      return count;
    };
  }
}
