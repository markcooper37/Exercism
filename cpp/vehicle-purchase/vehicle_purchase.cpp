#include "vehicle_purchase.h"

namespace vehicle_purchase {

    // needs_license determines whether a license is needed to drive a type of vehicle.
    bool needs_license(std::string kind){
        return kind == "car" || kind == "truck";
    }

    // choose_vehicle recommends a vehicle for selection.
    std::string choose_vehicle(std::string option1, std::string option2) {
        if (option1 < option2) return option1 + " is clearly the better choice.";
        else return option2 + " is clearly the better choice.";
    }

    // calculate_resell_price calculates how much a vehicle can resell for at a certain age.
    double calculate_resell_price(double original_price, double age) {
        if (age < 3) return original_price * 0.8;
        else if (age < 10) return original_price * 0.7;
        else return original_price * 0.5;
    }

}  // namespace vehicle_purchase