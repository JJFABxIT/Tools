#include <iostream>
#include <string>
#include <cstdlib>
#include <ctime>

std::string generate_password(int length) {
    const std::string characters =  "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
                                    "abcdefghijklmnopqrstuvwxyz"
                                    "0123456789!@£$%^&*";
    std::string password;
    srand(time(0));

    for (int i = 0; i < length; i++) {
        password += characters[rand() % characters.length()];
    }
    return password;
}

int main() {
    int length;

    std::cout << "Please enter the desired password length: ";
    std::cin >> length;

    std::string password = generate_password(length);
    std::cout << "Generated password = " <<  password << std::endl;

    return 0;
}