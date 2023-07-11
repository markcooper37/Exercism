#include <string>
#include <vector>

namespace election {

struct ElectionResult {
    std::string name{};
    int votes{};
};

int vote_count(ElectionResult& result) {
    return result.votes;
}

void increment_vote_count(ElectionResult& result, int votes) {
    result.votes += votes;
}

ElectionResult& determine_result(std::vector<ElectionResult>& final_count) {
    int winner{0};
    for (int i = 1; i < int(final_count.size()); i++) {
        if (final_count[i].votes > final_count[winner].votes) {
            winner = i;
        } 
    }
    final_count[winner].name = "President " + final_count[winner].name;
    return final_count[winner];
}

}  // namespace election