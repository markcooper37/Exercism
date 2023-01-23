export class List {
  constructor(values = []) {
    this.values = values;
  }

  append(list2) {
    let newValues = this.values;
    for (const value of list2.values) {
      newValues = [...newValues, value];
    };
    return new List(newValues);
  }

  concat(list) {
    let newList = new List(this.values);
    for (const value of list.values) {
      newList = newList.append(value);
    };
    return newList;
  }

  filter(predicate) {
    let newValues = [];
    for (const value of this.values) {
      if (predicate(value)) {
        newValues = [...newValues, value];
      };
    };
    return new List(newValues);
  }

  map(func) {
    let newValues = [];
    for (const value of this.values) {
      newValues = [...newValues, func(value)];
    };
    return new List(newValues);
  }

  length() {
    let count = 0;
    for (const value of this.values) {
      count++;
    };
    return count;
  }

  foldl(func, initial) {
    for (const value of this.values) {
      initial = func(initial, value);
    };
    return initial;
  }

  foldr(func, initial) {
    for (const value of this.reverse().values) {
      initial = func(initial, value);
    };
    return initial;
  }

  reverse() {
    let newValues = [];
    for (const value of this.values) {
      newValues = [value, ...newValues];
    };
    return new List(newValues);
  }
}
