class FlattenArray {
    static List flatten(List l) {
        def output = []
        for (item in l) {
            if (item instanceof List) {
                output.addAll(flatten(item).findAll {it != null})
            } else if (item != null) {
                output.add(item)
            }
        }
        return output
    }
}
