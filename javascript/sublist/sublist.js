export class List {
  constructor(list = []) {
    this.list = list;
  }

  compare(list2) {
    if (list2.list.length === this.list.length && isSublist(list2.list, this.list)) {
      return 'EQUAL';
    } else if (list2.list.length > this.list.length && isSublist(list2.list, this.list)) {
      return 'SUBLIST';
    } else if (isSublist(this.list, list2.list)) {
      return 'SUPERLIST';
    } else {
      return 'UNEQUAL';
    };
  }
}

function isSublist(list1, list2) {
  if (list2.length === 0) {
    return true;
  };
  for (let i = 0; i <= list1.length - list2.length; i++) {
    let sublist = true;
    for (let j = i; j < i + list2.length; j++) {
      if (list1[j] !== list2[j - i]) {
        sublist = false;
        break;
      };
    };
    if (sublist) {
      return true;
    };
  };
  return false;
}
