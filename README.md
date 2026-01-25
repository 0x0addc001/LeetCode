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

### 20260107

- [437. Path Sum III](https://leetcode.cn/problems/path-sum-iii/)

  ```python
  # PrefixSum 前缀和 O(n)
  class Solution:
      def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
          # prefixsum
          prefixdict=collections.defaultdict(int)
          
          # tot or diff == tarsum
          def dfs(node, legacy):
              if not node:
                  return 0
              res=0
              cur=legacy+node.val
              # add self res
              # 1. tot
              if cur==targetSum:
                  res+=1
              # 2. diff
              # root -> ...... -> node == cur
              # x -> ... -> node == tarsum
              # root -> ... -> pre_x == cur-tarsum
              res+=prefixdict[cur-targetSum]
  
              # add left res and right res
              prefixdict[cur]+=1
              res+=dfs(node.left, cur)
              res+=dfs(node.right, cur)
              prefixdict[cur]-=1 # +1-1 exclusive prefix to this branch
              return res
  
          return dfs(root, 0)
  ```
  

### 20260108

- [994. Rotting Oranges](https://leetcode.cn/problems/rotting-oranges/)

  ```python
  class Solution:
      def orangesRotting(self, grid: List[List[int]]) -> int:
          # enque
          # bfs
          # check 
          if not grid:
              return 0
          m=len(grid)
          n=len(grid[0])
          ctr=0
          que=deque()
          for i in range(m):
              for j in range(n):
                  if grid[i][j]==1:
                      ctr+=1
                  if grid[i][j]==2:
                      que.append((i,j,0))
          if ctr==0:
              return 0
          minutes=0
          while que:
              x,y,minutes=que.popleft()
              for dx,dy in [(1,0),(-1,0),(0,1),(0,-1)]:
                  if 0<=x+dx<m and 0<=y+dy<n and grid[x+dx][y+dy]==1:
                      grid[x+dx][y+dy]=2
                      que.append((x+dx,y+dy,minutes+1))
                      ctr-=1
          if ctr==0:
              return minutes
          else:
              return -1
  ```

### 20260115

- [78. Subsets](https://leetcode.cn/problems/subsets/)

  ```python
  # Bit 二进制
  class Solution:
      def subsets(self, nums: List[int]) -> List[List[int]]:
          ans=[]
          n=len(nums)
          for i in range(2**n):
              res=[]
              b=bin(i)[2:].zfill(n)
              for j in range(n):
                  if b[j]=='1':
                      res.append(nums[j])
              ans.append(res)
          return ans
  ```

### 20260116

- [78. Subsets](https://leetcode.cn/problems/subsets/)

  ```python
  # Backtracking 回溯
  class Solution:
      def subsets(self, nums: List[int]) -> List[List[int]]:
          ans=[[]]
          # 1,2,3
          # longest paths: 1->2->3, 2->3, 3
          # always look forward so that there's not repetition
          def dfs(root,unseens):
              if not unseens:
                  return
              for i,unseen in enumerate(unseens):
                  if root:
                      node=root+[unseen]
                  else:
                      node=[unseen]
                  ans.append(node)
                  dfs(node,unseens[i+1:])
          dfs([],nums)
          return ans
  ```
  

### 20260116

- [131. Palindrome Partitioning](https://leetcode.cn/problems/palindrome-partitioning/)

  ```python
  # Backtracking 回溯
  class Solution:
      def partition(self, s: str) -> List[List[str]]:
          # a a b
          # aab
          if not s:
              return []
          res=[]
          def is_palindrome(st):
              return st==st[::-1]
          def dfs(start,history):
              if start==len(s): # end
                  res.append(history[:]) # [:] for deep copy!!!
                  return
              for end in range(start+1,len(s)+1):
                  if is_palindrome(s[start:end]):
                      history.append(s[start:end])
                      dfs(end,history)
                      history.pop()
          dfs(0,[])
          return res
  ```

  ```python
  # Backtracking 回溯 + DP Optimization
  class Solution:
      def partition(self, s: str) -> List[List[str]]:
          # a a b
          # aab
          if not s:
              return []
          res=[]
          # def is_palindrome(st):
          #     return st==st[::-1]
          is_palindrome=[[False] * len(s) for _ in range(len(s))]
          for i in range(len(s)): # 1
              is_palindrome[i][i]=True
          for i in range(len(s)-1): # 2
              # is_palindrome[i][i+1]=(s[i:i+2]==s[i+1:i-1:-1]
              is_palindrome[i][i+1]=(s[i]==s[i+1])
          for length in range(3,len(s)+1): # >2
              for i in range(len(s)-length+1):
                  # is_palindrome[i][i+length-1]=(s[i:i+length]==s[i+length-1:i-1:-1]
                  is_palindrome[i][i+length-1]=(s[i]==s[i+length-1]) and is_palindrome[i+1][i+length-1-1]
  
          def dfs(start,history):
              if start==len(s): # end
                  res.append(history[:]) # [:] for deep copy!!!
                  return
              for end in range(start,len(s)):
                  # if is_palindrome(s[start:end]):
                  if is_palindrome[start][end]:
                      history.append(s[start:end+1])
                      dfs(end+1,history)
                      history.pop()
          dfs(0,[])
          return res
  ```

  
