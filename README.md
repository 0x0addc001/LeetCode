# LeetCode

## 2026

### 20260102

- [240. Search a 2D Matrix II](https://leetcode.cn/problems/search-a-2d-matrix-ii/)

  ```python
  class Solution:
      def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
          m=len(matrix)
          n=len(matrix[0])
          i=0
          j=n-1
          while i<m and j>=0:
              if matrix[i][j]==target:
                  return True
              elif matrix[i][j]>target:
                  j-=1 # left dicrease
              else:
                  i+=1 # down increase
          return False
  ```

### 20260103

- [24. Swap Nodes in Pairs](https://leetcode.cn/problems/swap-nodes-in-pairs/)

  ```python
  class Solution:
      def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
          # end case
          if not head or not head.next:
              return head
          nex=head.next
          head.next=self.swapPairs(nex.next)
          nex.next=head
          return nex
  ```


### 20260104

- [41. First Missing Positive](https://leetcode.cn/problems/first-missing-positive/)

  ```python
  class Solution:
      def firstMissingPositive(self, nums: List[int]) -> int:
          # sort
          # check
          nums.sort()
          ans=0
          for i in range(len(nums)):
              if nums[i]>0:
                  if nums[i]==ans+1:
                      ans+=1
                  elif nums[i]>ans+1:
                      return ans+1
          return ans+1
  ```

### 20260105

- [437. Path Sum III](https://leetcode.cn/problems/path-sum-iii/)

  ```python
  # dfs O(n^2)
  class Solution:
      def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
          # downward path -> dfs -> recur
          if not root:
              return 0
          ctr=0
          # start from root to somewhere
          ctr+=self.rootSum(root,targetSum)
          ctr+=self.pathSum(root.left,targetSum)
          ctr+=self.pathSum(root.right,targetSum)
          return ctr
  
      def rootSum(self, root: Optional[TreeNode], targetSum: int) -> int:
          if not root:
              return 0
          ctr=0
          if root.val==targetSum:
              ctr+=1
          ctr+=self.rootSum(root.left,targetSum-root.val)
          ctr+=self.rootSum(root.right,targetSum-root.val)
          return ctr
  ```

  ```python
  # 前缀和 O(n)
  ```

  
