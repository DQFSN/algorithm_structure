#include<vector>
#include<iostream>

using namespace std;

class Solution {
public:
    vector<vector<int>>  transpose(vector<vector<int>>& A) {

        // solution 1
        // if (!A.size()) return A;
        // vector<vector<int>> B(A[0].size());
        // int rnum = A.size();
        // int cnum = A[0].size();
        // for (auto c = 0; c < cnum; c++)
        // {
        //     for (auto r = 0; r < rnum; r++)
        //     {
        //             B[c].push_back(A[r][c]);
        //     }
        // }
        
        //solution 2

        vector<vector<int>> B(A[0].size(), vector<int>(A.size(),0));

        for (int i = 0; i < A.size(); i++)
        {
            for (int j = 0; j <A[0].size(); j++)
            {
                B[j][i] = A[i][j];
            }
        }

        return B;
    } 
};


int main() {
    // vector<vector<int>> A {{1,2,3}, {4,5,6},{7,8,9}};
    vector<vector<int>> A {{1,2,3}, {4,5,6}};
    vector<vector<int>> B = Solution().transpose(A);

    for (auto &&r : B)
    {
        for (auto &&v : r)
        {
            cout<<v<<", ";
        }
        
        cout<<endl;
    }
    

    return 0;
}