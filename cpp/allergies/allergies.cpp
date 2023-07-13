#include "allergies.h"
#include <string>
#include <unordered_set>
#include <map>

namespace allergies {
    bool allergy_test::is_allergic_to(std::string allergen) {
        return allergens[allergen] & score;
    }

    std::unordered_set<std::string> allergy_test::get_allergies() {
        std::unordered_set<std::string> allergies {};
        for (auto &[allergen, allergen_score] : allergens) {
            if (allergen_score & score) {
                allergies.insert(allergen);
            }
        }
        return allergies;
    }
}  // namespace allergies
