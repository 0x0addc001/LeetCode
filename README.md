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
  

### 20260125

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


### 20260126

- [51. N-Queens](https://leetcode.cn/problems/n-queens/)

  ```python
  # Backtracking 回溯
  class Solution:
      def solveNQueens(self, n: int) -> List[List[str]]:
          res=[]
          def is_safe(line,i,q_list):
              # ..Q..
              # .xxx.
              # x.x.x
              for q_ln,q_col in q_list:
                  if i==q_col:
                      return False
                  # \
                  # i+line-q_col
                  right=q_col+line-q_ln
                  if i==right:
                      return False
                  # /
                  # i-line+q_col
                  left=q_col-line+q_ln
                  if i==left:
                      return False
              return True
  
          def dfs(line,history,q_list):
              if line==n:
                  res.append(history[:])
                  return
              for i in range(n):
                  if is_safe(line,i,q_list):
                      record='.'*i+'Q'+'.'*(n-i-1)
                      history.append(record)
                      q_list.append((line,i))
                      dfs(line+1,history,q_list)
                      history.pop()
                      q_list.pop()
          dfs(0,[],[])
          return res
  ```


### 20260201

- [347. Top K Frequent Elements](https://leetcode.cn/problems/top-k-frequent-elements/)

  ```python
  # heapq.nlargest 堆
  class Solution:
      def topKFrequent(self, nums: List[int], k: int) -> List[int]:
          freq_dict=defaultdict(int)
          for i in nums:
              freq_dict[i]+=1
          freq_list=[]
          for key,val in freq_dict.items():
              freq_list.append((key,val))
          # freq_list=sorted(freq_list,key=lambda x:x[1],reverse=True)
          # res=[]
          # for i in range(k):
          #     res.append(freq_list[i][0])
          res=[key for key,val in heapq.nlargest(k,freq_list,key=lambda x:x[1])]
          return res
  ```

  ```python
  # Counter
  class Solution:
      def topKFrequent(self, nums: List[int], k: int) -> List[int]:
          freq=Counter(nums)
          # return [key for key,val in heapq.nlargest(k,freq.items(),key=lambda x:x[1])]
          return [key for key,val in freq.most_common(k)]
  ```


### 20260210

- [279. Perfect Squares](https://leetcode.cn/problems/perfect-squares/)

  ```python
  # dp
  class Solution:
      def numSquares(self, n: int) -> int:
          # with all the subs plusing sq
          # accept the small
  
          # init
          dp=[float('inf') for _ in range(n+1)]
          dp[0]=0
          # prep
          sqlist=[i**2 for i in range(int(n**0.5)+1)] # +1
          # calc
          for i in range(1,n+1):
              
              # for j in range(n): # from 0
              #     for k in range(1,i-j):
              #         if k**2==i-j: # just eq
              #             dp[i]=min(dp[i],1+dp[j])
  
              # for k in range(1,int(i**0.5)+1):
              #     dp[i]=min(dp[i],1+dp[i-k**2])
  
              for sq in sqlist:
                  if sq>i:
                      break
                  dp[i]=min(dp[i],1+dp[i-sq])
  
          return dp[n]
  ```


### 20260211

- [1143. Longest Common Subsequence](https://leetcode.cn/problems/longest-common-subsequence/)

  ```python
  class Solution:
      def longestCommonSubsequence(self, text1: str, text2: str) -> int:
          m=len(text1)
          n=len(text2)
          '''
          dp=[[0 for _ in range(n+1)] for _ in range(m+1)] # +1 to consider empty str and universality of trasnformation function
          # \ 0 1 2
          # 0 0 
          # 1
          # 2
          for i in range(1,m+1):
              for j in range(1,n+1):
                  # ==
                  if text1[i-1]==text2[j-1]:
                      dp[i][j]=dp[i-1][j-1]+1
                  # t1 has sth more: adbc abc
                  # t2 has sth more: abc adbc
                  # just diff: abc adc
                  elif text1[i-1]!=text2[j-1]:
                      # dp[i][j]=max(dp[i-1][j-1],dp[i][j-1],dp[i-1][j])
                      dp[i][j]=max(dp[i][j-1],dp[i-1][j])
          return dp[m][n]
          '''
          # O(n) linear optimization: for each time we only use the last and cur row of the dp table
          pre=[0 for _ in range(n+1)]
          for i in range(1,m+1):
              cur=[0 for _ in range(n+1)]
              for j in range(1,n+1):
                  if text1[i-1]==text2[j-1]:
                      cur[j]=pre[j-1]+1
                  else:
                      cur[j]=max(pre[j],cur[j-1])
              pre=cur
          return pre[n]
  ```


### 20260213

- [394. Decode String](https://leetcode.cn/problems/decode-string/)

  ```python
  class Solution:
      def decodeString(self, s: str) -> str:
          # [
          # ]
          n=len(s)
          digits_new=0
          letters_new=''
          stack=[]
          for c in s:
              if c.isdigit(): # 数字
                  digits_new=digits_new*10+int(c)
              elif c=='[': # 过去压栈
                  stack.append((letters_new,digits_new))
                  letters_new=''
                  digits_new=0
              elif c==']': # 当前清算
                  letters_old,digits_old=stack.pop()
                  letters_new=letters_old+(letters_new*digits_old)
              else: # 字母
                  letters_new+=c
          return letters_new
  ```


### 20260216

- [739. Daily Temperatures](https://leetcode.cn/problems/daily-temperatures/)

  ```python
  class Solution:
      def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
          n=len(temperatures)
          res=[0 for _ in range(n)]
          # bf
          # --->
          # for i in range(n):
          #     for j in range(i,n):
          #         if temperatures[j]>temperatures[i]:
          #             res[i]=j-i
          #             break
  
          # mono stack / monotonic stack
          # if we look all the way back from the future, then we don't need to wait for each future
          unsolved_stk=[]
          for i in range(n):
              # warmer day arrives then roll back all the cooler prev days
              while unsolved_stk and temperatures[i]>temperatures[unsolved_stk[-1]]:
                  prev=unsolved_stk.pop()
                  res[prev]=i-prev
              unsolved_stk.append(i)
          return res
  ```


### 20260218

- [84. Largest Rectangle in Histogram](https://leetcode.cn/problems/largest-rectangle-in-histogram/)

  ```python
  class Solution:
      def largestRectangleArea(self, heights: List[int]) -> int:
          # bf
          # n=len(heights)
          # res=[0 for _ in range(n)]
          # for i in range(n):
          #     lowest=float('inf')
          #     for j in range(i,-1,-1):
          #         lowest=min(lowest,heights[j])
          #         res[i]=max(res[i],lowest*(i-j+1))
          # return max(res)
  
          # mono stack
          # could we save a cache for each position which means the value of it and its past, so that we don't need to recompute it again?
          # it is particularly important when it comes to the bottleneck like '1', why? because it tightly control the previous information (even such kind of number is in the middle is still needed for reconsideration), thus it can suggest us whether to get (lowest*(i-prev_start+1)) or forget (new_area)
          # so that mono stk is a decreasing stk
          # '1' and '2' give the information of a period
          # n=len(heights)
          # res=0
          # mono_stk=[] # h[i+1]<=h[i] # (start,end,lowest)
          # for i in range(n):
          #     j=len(mono_stk)-1
          #     global_lowest=float('inf')
          #     while mono_stk and heights[i]>mono_stk[-1][2] and j>=0:
          #         # 有对手&打得过
          #         # start,end,lowest=mono_stk.pop()
          #         # mono_stk.append((start,end,lowest))
          #         start,end,lowest=mono_stk[j]
          #         global_lowest=min(global_lowest,lowest)
          #         res=max(res,heights[i]*1,global_lowest*(i-start+1))
          #         j-=1
          #         if j<0:
          #             break
          #     else: # <= or no mono_stack
          #         # 打不过就加入
          #         # if mono_stk and heights[i]<=mono_stk[-1][2]:
          #         if mono_stk:
          #             start,end,lowest=mono_stk.pop()
          #             end+=1
          #             lowest=heights[i]
          #             mono_stk.append((start,end,lowest))
          #             # res=max(res,lowest*(i-start+1))
          #             global_lowest=min(global_lowest,lowest)
          #             res=max(res,global_lowest*(i-start+1))
          #             while mono_stk and j>=0:
          #                 prev_start,prev_end,prev_lowest=mono_stk[j]
          #                 global_lowest=min(global_lowest,prev_lowest)
          #                 res=max(res,global_lowest*(i-prev_start+1))
          #                 j-=1
          #             continue
          #     #     # 没对手就立派
          #     #     else: # no mono_stack
          #     #         mono_stk.append((i,i,heights[i]))
          #     #         res=max(res,heights[i]*1)
          #     #         continue
          #     mono_stk.append((i,i,heights[i]))
          #     res=max(res,heights[i]*1)
          # return res
  
          # mono stack
          n=len(heights)
          res=0
          mono_stk=[] # increasing mono_stk
          heights.append(0) # add a 垃圾回收车/结尾小推车
          for i in range(n+1):
              while mono_stk and heights[i]<heights[mono_stk[-1]]: # the right boundary can't lower the bar for those talls, nor the left boundary, or it will be the same question (downgrading the question) as if they were the same bar as they are: [2,3,3,2]->[3,3]/[2,2,2,2]; [2,3,3,2] is now equivalent to [2,2,2,2]
              # Handle the tallers at first, then the shorters, 先处理高子，再处理矮子
                  prev=mono_stk.pop()
                  height=heights[prev] # popped height as lowest, cuz the mono_stk is decreasing backward
                  # width=(i-1)-0+1 if not mono_stk else (i-1)-(mono_stk[-1]+1)+1
                  width=i if not mono_stk else i-1-mono_stk[-1]
                  res=max(res,height*width)
              # for all
              mono_stk.append(i)
          return res
  ```


### 20260219

- [152. Maximum Product Subarray](https://leetcode.cn/problems/maximum-product-subarray/)

  ```python
  class Solution:
      def maxProduct(self, nums: List[int]) -> int:
          n=len(nums)
          # i: start from i
          # j: end at j
          # _ 2 3 -2 4
          # 2
          # 3
          # -2
          # 4
  
          # T(n**2) S(n**2)
          # dp=[[0 for _ in range(n)] for _ in range(n)]
          # res=-float('inf')
          # for i in range(n):
          #     for j in range(i,n):
          #         if i==j:
          #             dp[i][j]=nums[i]
          #         else: # j>i
          #             dp[i][j]=dp[i][j-1]*nums[j]
          #         res=max(res,dp[i][j])
  
          # T(n**2) S(n)
          # dp=[0 for _ in range(n)]
          # res=-float('inf')
          # for i in range(n):
          #     for j in range(i,n):
          #         if i==j:
          #             dp[j]=nums[j]
          #         else: # j>i
          #             dp[j]=dp[j-1]*nums[j]
          #         res=max(res,dp[j])
  
          # T(n) S(1)
          res=nums[0]
          acc_max=nums[0]
          acc_min=nums[0] # used for negative flipping cases
          for i in range(1,n):
              pre_max=acc_max
              pre_min=acc_min
              acc_max=max(nums[i],nums[i]*pre_max,nums[i]*pre_min)
              acc_min=min(nums[i],nums[i]*pre_max,nums[i]*pre_min)
              res=max(res,acc_max) # its not necessary to involve the current num
          return res
  ```


### 20260220

- [416. Partition Equal Subset Sum](https://leetcode.cn/problems/partition-equal-subset-sum/)

  ```python
  # dp
  class Solution:
      def canPartition(self, nums: List[int]) -> bool:
          # 1 2 3 4
          # + - - +
          tot_sum=sum(nums)
          if tot_sum%2!=0: # odd
              return False
          half_sum=tot_sum//2
          # we hope to have a subset which is equals to half_sum
          # n_sum problem
          # def n_sum(given_nums,target):
          #     given_nums_len=len(given_nums)
          #     if given_nums_len==0:
          #         return target==0
          #     if given_nums_len==1:
          #         return target==given_nums[0]
          #     for i in range(given_nums_len):
          #         if n_sum(nums[:i]+nums[i+1:],half_sum-given_nums[i]):
          #             return True
          #     return False
          # return n_sum(nums,half_sum)
          sum_set=set([0]) # [0] for standalone num choosing
          for num in nums:
              new_sum_set=set()
              for s in sum_set:
                  res=num+s
                  if res==half_sum:
                      return True
                  elif res<half_sum:
                      # sum_set.add(res) # Set changed size during iteration
                      new_sum_set.add(res)
              sum_set |= new_sum_set # union in place
          return False
  ```

  
