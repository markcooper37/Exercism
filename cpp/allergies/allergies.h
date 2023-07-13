#if !defined(ALLERGIES_H)
#define ALLERGIES_H
#include <unordered_set>
#include <map>

namespace allergies {
    class allergy_test {
        int score;
        std::map<std::string, int> allergens = {{"eggs",1},{"peanuts",2},{"shellfish",4},{"strawberries",8},{"tomatoes",16},{"chocolate",32},{"pollen",64},{"cats",128}};
    public:
        allergy_test(int num) {score = num;};
        bool is_allergic_to(std::string allergen);
        std::unordered_set<std::string> get_allergies();
    };
}  // namespace allergies

#endif // ALLERGIES_H