# LeetCode

## 20260102

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

## 20260103

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


## 20260104

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

## 20260105

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

## 20260107

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
  

## 20260108

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

## 20260115

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

## 20260116

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
  

## 20260125

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


## 20260126

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


## 20260201

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


## 20260210

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


## 20260211

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


## 20260213

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


## 20260216

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


## 20260218

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


## 20260219

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


## 20260220

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


## 20260221

- [32. Longest Valid Parentheses](https://leetcode.cn/problems/longest-valid-parentheses/)

  ```python
  # bf
  class Solution:
      def longestValidParentheses(self, s: str) -> int:
          # one pass but not necessarily continuous
          # ctr=0
          # stk=[]
          # for c in s:
          #     if c=='(':
          #         stk.append('(')
          #     elif c==')':
          #         if stk:
          #             stk.pop()
          #             ctr+=2
          # return ctr
          # one pass and continuous
          cont_matched_stk=[]
          par_stk=[]
          pause_flag=False
          for c in s:
              if c=='(':
                  par_stk.append('(')
              elif c==')':
                  if par_stk and par_stk[-1]=='(':
                      par_stk.pop()
                      # check swallow then check merge:
                      if cont_matched_stk and cont_matched_stk[-1][1]==len(par_stk)+1: # only make friends with equal or higher level
                          num,lev=cont_matched_stk.pop()
                          num+=2 # swallow smaller
                          if cont_matched_stk and cont_matched_stk[-1][1]==len(par_stk):
                              num_,lev_=cont_matched_stk.pop()
                              num+=num_ # merge equal
                          cont_matched_stk.append((num,len(par_stk)))
                      # check merge
                      elif cont_matched_stk and cont_matched_stk[-1][1]==len(par_stk):
                          num,lev_=cont_matched_stk.pop()
                          num+=2 # merge equal
                          cont_matched_stk.append((num,len(par_stk)))
                      else:
                          cont_matched_stk.append((2,len(par_stk)))
                  else:
                      # unexpected
                      par_stk=[]
                      cont_matched_stk.append((0,0))
          res=max(cont_matched_stk, key=lambda x:x[0]) if cont_matched_stk else (0,0)
          return res[0]
  ```

  ```python
  class Solution:
      def longestValidParentheses(self, s: str) -> int:
          left=right=0
          res=0
          for c in s:
              if c=='(':
                  left+=1
              else:
                  right+=1
              if left==right:
                  res=max(res,left*2)
              elif right>left:
                  left=right=0
          left=right=0
          for c in reversed(s): # (() -> )((
              if c=='(':
                  left+=1
              else:
                  right+=1
              if left==right:
                  res=max(res,left*2)
              elif left>right:
                  left=right=0
          return res
  ```


## 20260222

- [75. Sort Colors](https://leetcode.cn/problems/sort-colors/)

  ```python
  class Solution:
      def sortColors(self, nums: List[int]) -> None:
          """
          Do not return anything, modify nums in-place instead.
          """
          # nums.sort()
          
          # T(n**2) two pass	
          # freq=Counter(nums)
          # idx=0
          # for i in range(3):
          #     for j in range(freq[i]):
          #         nums[idx]=i
          #         idx+=1
  
          # T(n) one pass
          ctr0=ctr1=ctr2=0
          for num in nums:
              if num==0:
                  nums[ctr2]=2 # 往后骚
                  ctr2+=1
                  nums[ctr1]=1 # 往后骚
                  ctr1+=1
                  nums[ctr0]=0
                  ctr0+=1
              elif num==1:
                  nums[ctr2]=2 # 往后骚
                  ctr2+=1
                  nums[ctr1]=1
                  ctr1+=1
              elif num==2:
                  nums[ctr2]=2
                  ctr2+=1
  ```


## 20260311

- [31. Next Permutation](https://leetcode.cn/problems/next-permutation/)

  ```python
  class Solution:
      def nextPermutation(self, nums: List[int]) -> None:
          """
          Do not return anything, modify nums in-place instead.
          """
          # O(n**2)
          # 1234
          # 1243 -swap-> 1 342 -sort->13 24
          # 1324
          # 1342
          # 1423
          # 1432
          # 2134
          # ...
          # 4321
          n=len(nums)
          if n==1:
              return nums
          for i in range(n-2,-1,-1): # back
              # min greater
              min_greater_num=float('inf')
              min_greater_idx=-1
              for j in range(i+1,n):
                  if nums[j]>nums[i] and nums[j]<min_greater_num:
                      min_greater_num=nums[j]
                      min_greater_idx=j
              if min_greater_idx!=-1 and nums[min_greater_idx]>nums[i]:
                  # swap
                  nums[i],nums[min_greater_idx]=nums[min_greater_idx],nums[i]
                  # sort
                  nums[i+1:]=sorted(nums[i+1:])
                  return
          nums.reverse()
          
          
          
          # O(n) optimization
          n=len(nums)
          i=n-2
          # find the first adjacent ascend, otherwise is descend
          while i>=0 and nums[i]>=nums[i+1]:
              i-=1
          if i>=0:
              j=n-1
              # find the first num greater than nums[i] (min greater) along the ascend list (descend in reverse), the worst case is i+1
              while nums[j]<=nums[i]:
                  j-=1
              nums[i],nums[j]=nums[j],nums[i]
          nums[i+1:]=nums[i+1:][::-1]
  ```


## 20260312

- [287. Find the Duplicate Number](https://leetcode.cn/problems/find-the-duplicate-number/)

  ```python
  class Solution:
      def findDuplicate(self, nums: List[int]) -> int:
          n=len(nums)
          
          # for i in range(n-1):
          #     for j in range(i+1,n):
          #         if nums[i]==nums[j]:
          #             return nums[i]
          # return -1
  
          # nums.sort()
          # for i in range(n-1):
          #     if nums[i]==nums[i+1]:
          #         return nums[i]
          # return -1
          
          kvcache=defaultdict(int)
          for i in nums:
              if kvcache[i]>0:
                  return i
              kvcache[i]+=1
          return -1
  ```
  

## 20260314

- [198. House Robber](https://leetcode.cn/problems/house-robber/)

  ```python
  class Solution:
      def rob(self, nums: List[int]) -> int:
          n=len(nums)
          if n==1:
              return nums[0]
              
          # S(n)
          # benefit=[0 for _ in range(n)]
          # benefit[0]=nums[0]
          # benefit[1]=max(nums[0],nums[1])
          # for i in range(2,n):
          #     benefit[i]=max(benefit[i-2]+nums[i],benefit[i-1])
          # return benefit[n-1]
  
          # S(1)
          pre=nums[0]
          cur=max(nums[0],nums[1])
          for i in range(2,n):
              pre,cur=cur,max(pre+nums[i],cur)
          return cur
  ```


## 20260318

- [148. Sort List](https://leetcode.cn/problems/sort-list/)

  ```python
  # Definition for singly-linked list.
  # class ListNode:
  #     def __init__(self, val=0, next=None):
  #         self.val = val
  #         self.next = next
  class Solution:
      def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
          if not head or not head.next: # at least 2
              return head
          fast=head.next # !!!!! designed for 2-node situation
          slow=head
          while fast and fast.next:
              fast=fast.next.next
              slow=slow.next
          mid=slow.next
          slow.next=None # !!!!! cut off relationship between the 2
          list1=self.sortList(head)
          list2=self.sortList(mid)
          return self.merge(list1,list2)
  
      def merge(self,head1,head2): # !!!!! self
          dummy=ListNode()
          temp=dummy
          while head1 and head2:
              if head1.val<=head2.val:
                  temp.next=head1
                  head1=head1.next
                  temp=temp.next
              else:
                  temp.next=head2
                  head2=head2.next
                  temp=temp.next
          if head1:
              temp.next=head1
          if head2:
              temp.next=head2
          return dummy.next
  ```


## 20260321

- [146. LRU Cache](https://leetcode.cn/problems/lru-cache/)

  ```python
  class LRUCache:
  
      def __init__(self, capacity: int):
          self.capacity=capacity
          self.cnt=0
          self.kvcache=OrderedDict()
  
      def get(self, key: int) -> int:
          if key in self.kvcache:
              value=self.kvcache[key]
              self.kvcache.move_to_end(key)
              return value
          else:
              return -1
          
      def put(self, key: int, value: int) -> None:
          if key in self.kvcache:
              self.kvcache[key]=value
              self.kvcache.move_to_end(key)
          else:
              if self.cnt>=self.capacity:
                  # evict
                  self.kvcache.popitem(last=False)
                  self.cnt-=1
              # add
              self.kvcache[key]=value
              self.cnt+=1
          
  
  
  # Your LRUCache object will be instantiated and called as such:
  # obj = LRUCache(capacity)
  # param_1 = obj.get(key)
  # obj.put(key,value)
  ```

  

## 20260324

- [3. Longest Substring Without Repeating Characters](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)

  ```python
  class Solution:
      def lengthOfLongestSubstring(self, s: str) -> int:
          n=len(s)
          cnt=0
          seen=defaultdict(int)
          left=0
          for right in range(n):
              while left<=right and seen[s[right]]>0:
                  seen[s[left]]-=1
                  left+=1
              seen[s[right]]+=1
              cnt=max(cnt,right-left+1)
          return cnt
  ```

- [25. Reverse Nodes in k-Group](https://leetcode.cn/problems/reverse-nodes-in-k-group/)

  ```python
  # Definition for singly-linked list.
  # class ListNode:
  #     def __init__(self, val=0, next=None):
  #         self.val = val
  #         self.next = next
  class Solution:
      def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
          # find pre
          # find post
          # reverse pre [...] post
          # reconnect
          # move pre post
  
          # X->A->B->C->D
  
          # X->B->C->D
          # Y->A
  
          # X->C->D
          # Y->B->A
          def reverse(hd):
              if not hd:
                  return hd
              dmy=ListNode()
              dmy2=ListNode()
              dmy.next=hd
              while dmy and dmy.next:
                  dmy2next=dmy2.next
                  dmy2.next=dmy.next
                  dmy.next=dmy.next.next
                  dmy2.next.next=dmy2next
              return dmy2.next
  
          def reverse2(hd):
              if not hd:
                  return hd
              prev=None
              curr=hd
              while curr: # iter point-back
                  currnext=curr.next
                  curr.next=prev
                  prev=curr
                  curr=currnext
              return prev
  
          dummy=ListNode()
          dummy.next=head
          pre=post=dummy
          while pre.next:
              cnt=k
              post=pre
              while cnt>0 and post.next: # the last one of (this group or the entire list)
                  post=post.next
                  cnt-=1
              if cnt>0:
                  break
              postnext=post.next
              post.next=None # cut
              newList=reverse(pre.next)
              pre.next=newList
              while pre and pre.next:
                  pre=pre.next
              pre.next=postnext
          return dummy.next
  ```


## 20260403

- [206. Reverse Linked List](https://leetcode.cn/problems/reverse-linked-list/)

  ```python
  # Definition for singly-linked list.
  # class ListNode:
  #     def __init__(self, val=0, next=None):
  #         self.val = val
  #         self.next = next
  class Solution:
      def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
          # 1 <- 2 3 4 5
          if not head:
              return head
          '''
          pre=head
          cur=head.next
          pre.next=None # or will sink into dead recur between 1 & 2
          while pre and cur:
              temp=cur.next # 3
              cur.next=pre # 2->1
              pre=cur # 1=>2
              cur=temp # 2=>3
          return pre
          '''
          # N<-1 2
          pre=None
          cur=head
          while cur:
              tp=cur.next # 2
              cur.next=pre # 1->N
              pre=cur # N=>1
              cur=tp # 1=>2
          return pre
  ```

- [23. Merge k Sorted Lists](https://leetcode.cn/problems/merge-k-sorted-lists/)

  ```python
  # Definition for singly-linked list.
  # class ListNode:
  #     def __init__(self, val=0, next=None):
  #         self.val = val
  #         self.next = next
  class Solution:
      def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
          # bf
          if not lists:
              return None
          dummy=ListNode()
          cur_ptr=dummy
          n=len(lists)
          while True:
              min_val=float('inf')
              min_idx=-1
              for i in range(n):
                  node=lists[i]
                  if node and node.val<min_val:
                      min_val=node.val
                      min_idx=i
              if min_idx==-1:
                  break
              cur_ptr.next=lists[min_idx]
              cur_ptr=cur_ptr.next
              print(cur_ptr.val)
              # min_ptr=min_ptr.next # only update local value, doesn't update the value in the original list
              # for node in lists: # it doesn't work too
              #     if node.val==min_ptr.val:
              #         node=node.next
              #         break
              lists[min_idx]=lists[min_idx].next
          return dummy.next
          # optimize1: lists -> heap o(nlogk)
          # optimize2: merge 2 list for each loop o(nlogk)
  ```

## 20260405
- [23. Merge k Sorted Lists](https://leetcode.cn/problems/merge-k-sorted-lists/)

   ```python
    # Definition for singly-linked list.
    # class ListNode:
    #     def __init__(self, val=0, next=None):
    #         self.val = val
    #         self.next = next
    class Solution:
        def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
            # bf
            if not lists:
                return None
            dummy=ListNode()
            cur_ptr=dummy
            n=len(lists)
            while True:
                min_val=float('inf')
                min_idx=-1
                for i in range(n):
                    node=lists[i]
                    if node and node.val<min_val:
                        min_val=node.val
                        min_idx=i
                if min_idx==-1:
                    break
                cur_ptr.next=lists[min_idx]
                cur_ptr=cur_ptr.next
                print(cur_ptr.val)
                # min_ptr=min_ptr.next # only update local value, doesn't update the value in the original list
                # for node in lists: # it doesn't work too
                #     if node.val==min_ptr.val:
                #         node=node.next
                #         break
                lists[min_idx]=lists[min_idx].next
            return dummy.next
    
            # optimize1: lists -> heap o(nlogk) [priority q]
            if not lists:
                return None
            dummy=ListNode()
            cur=dummy
            hp=[]
            # for node in lists:
            #     if node:
            #         heapq.heappush(hp,(node.val,id(node),node)) # avoid "TypeError: '<' not supported between instances of 'Node' and 'Node'", because if val=val, then move to the next criteria
            # cleaner impl
            hp=[(node.val,id(node),node) for node in lists if node]
            heapq.heapify(hp)
            while hp:
                val,idx,nd=heapq.heappop(hp)
                cur.next=nd
                cur=cur.next
                if nd.next:
                    heapq.heappush(hp,(nd.next.val,id(nd.next),nd.next))
            return dummy.next
            
            # bf2
            if not lists:
                return None
            def merge2(n1,n2):
                if not n1:
                    return n2
                if not n2:
                    return n1
                dmy=ListNode()
                cur=dmy
                while n1 and n2:
                    if n1.val<n2.val:
                        cur.next=n1
                        n1=n1.next
                    else:
                        cur.next=n2
                        n2=n2.next
                    cur=cur.next
                cur.next=n1 if n1 else n2
                return dmy.next
            res=None
            for node in lists:
                res=merge2(res,node)
            return res
            
            # optimize2: merge 2 list for each loop o(nlogk)
            if not lists:
                return None
            def merge2(n1,n2):
                if not n1:
                    return n2
                if not n2:
                    return n1
                dmy=ListNode()
                cur=dmy
                while n1 and n2:
                    if n1.val<n2.val:
                        cur.next=n1
                        n1=n1.next
                    else:
                        cur.next=n2
                        n2=n2.next
                    cur=cur.next
                cur.next=n1 if n1 else n2
                return dmy.next
            # DC Tree
            while len(lists)>1: # from 2**n to 1
                merged_lists=[]
                for i in range(0,len(lists),2): # merge 2 by 2
                    node1=lists[i]
                    node2=lists[i+1] if (i+1) < len(lists) else None
                    merged_lists.append((merge2(node1,node2)))
                lists=merged_lists
            return lists[0]
   ```

- [21. Merge Two Sorted Lists](https://leetcode.cn/problems/merge-two-sorted-lists/)

  ```python
  # Definition for singly-linked list.
  # class ListNode:
  #     def __init__(self, val=0, next=None):
  #         self.val = val
  #         self.next = next
  class Solution:
      def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
          if not list1:
              return list2
          if not list2:
              return list1
          dummy=ListNode()
          cur=dummy
          while list1 and list2:
              if list1.val<list2.val:
                  cur.next=list1
                  list1=list1.next
              else:
                  cur.next=list2
                  list2=list2.next
              cur=cur.next
          cur.next=list1 if list1 else list2
          return dummy.next
  ```


## 20260406

- [141. Linked List Cycle](https://leetcode.cn/problems/linked-list-cycle/)

  ```python
  # Definition for singly-linked list.
  # class ListNode:
  #     def __init__(self, x):
  #         self.val = x
  #         self.next = None
  
  class Solution:
      def hasCycle(self, head: Optional[ListNode]) -> bool:
          if not head:
              return head
          slow=head
          fast=head.next
          while fast and fast.next:
              if fast==slow:
                  return True
              fast=fast.next.next
              slow=slow.next
          return False
  ```

- [2. Add Two Numbers](https://leetcode.cn/problems/add-two-numbers/)

  ```python
  # Definition for singly-linked list.
  # class ListNode:
  #     def __init__(self, val=0, next=None):
  #         self.val = val
  #         self.next = next
  class Solution:
      def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
          dummy=ListNode()
          cur=dummy
          overflow=0
          while l1 and l2:
              nex=ListNode()
              nex.val=((l1.val+l2.val+overflow)%10)
              overflow=((l1.val+l2.val+overflow)//10)
              cur.next=nex
              cur=cur.next
              l1=l1.next
              l2=l2.next
          while l1:
              nex=ListNode()
              nex.val=((l1.val+overflow)%10)
              overflow=((l1.val+overflow)//10)
              cur.next=nex
              cur=cur.next
              l1=l1.next
          while l2:
              nex=ListNode()
              nex.val=((l2.val+overflow)%10)
              overflow=((l2.val+overflow)//10)
              cur.next=nex
              cur=cur.next
              l2=l2.next
          if overflow==1:
              cur.next=ListNode(val=1)
          return dummy.next
  ```


## 20260407

- [19. Remove Nth Node From End of List](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)

  ```python
  # Definition for singly-linked list.
  # class ListNode:
  #     def __init__(self, val=0, next=None):
  #         self.val = val
  #         self.next = next
  class Solution:
      def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
          if not head:
              return head
          # prewalk k
          # set dummy
          # step num to end: cnt
          # step num to n before end: cnt-n+1
          # step num to n+1 brefore end: cnt-n+1-1=cnt-n
          cnt=0
          cur=head
          while cur:
              cnt+=1
              cur=cur.next
          dummy=ListNode(next=head)
          cur=dummy
          k=cnt-n
          while k>0:
              cur=cur.next
              k-=1
          cur.next=cur.next.next
          return dummy.next
  ```

- [138. Copy List with Random Pointer](https://leetcode.cn/problems/copy-list-with-random-pointer/)

  ```python
  """
  # Definition for a Node.
  class Node:
      def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
          self.val = int(x)
          self.next = next
          self.random = random
  """
  
  class Solution:
      def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
          # copy
          # mapping dict node_old:node_new
          dummy=Node(x=0)
          node_map=dict()
          cur=dummy
          node_old=head
          while node_old:
              node_new=Node(x=node_old.val,next=node_old.next,random=node_old.random)
              node_map[node_old]=node_new
              cur.next=node_new
              cur=cur.next
              node_old=node_old.next
          cur=dummy.next
          while cur:
              if cur.random:
                  cur.random=node_map[cur.random]
              cur=cur.next
          return dummy.next
  ```


## 20260408

- [208. Implement Trie (Prefix Tree)](https://leetcode.cn/problems/implement-trie-prefix-tree/)

  ```python
  # recur
  class Trie:
  
      def __init__(self):
          self.children=dict() # char:node
  
      def insert(self, word: str) -> None:
          if not word:
              if None not in self.children:
                  self.children[None]=None
          else:
              if word[0] in self.children: # word[0] is the char i nees to find next
                  self.children[word[0]].insert(word[1:])
              else:
                  temp=Trie()
                  temp.insert(word[1:])
                  self.children[word[0]]=temp
  
      def search(self, word: str) -> bool:
          if not word:
              if None not in self.children:
                  return False
              return True
          else:
              if word[0] in self.children:
                  return self.children[word[0]].search(word[1:])
              else:
                  return False
  
      def startsWith(self, prefix: str) -> bool:
          if not prefix:
              return True
          else:
              if prefix[0] in self.children:
                  return self.children[prefix[0]].startsWith(prefix[1:])
              else:
                  return False
  
  # iter
  class Trie:
  
      def __init__(self):
          self.children={} # char:node
          self.is_end=False
  
      def insert(self, word: str) -> None:
          node=self
          for char in word:
              if char not in node.children:
                  node.children[char]=Trie()
              node=node.children[char]
          node.is_end=True
  
      def search(self, word: str) -> bool:
          node=self
          for char in word:
              if char not in node.children:
                  return False
              node=node.children[char]
          return node.is_end
  
      def startsWith(self, prefix: str) -> bool:
          node=self
          for char in prefix:
              if char not in node.children:
                  return False
              node=node.children[char]
          return True
  
  # Your Trie object will be instantiated and called as such:
  # obj = Trie()
  # obj.insert(word)
  # param_2 = obj.search(word)
  # param_3 = obj.startsWith(prefix)
  ```


## 20260409

- [226. Invert Binary Tree](https://leetcode.cn/problems/invert-binary-tree/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, val=0, left=None, right=None):
  #         self.val = val
  #         self.left = left
  #         self.right = right
  class Solution:
      def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
          # standard recur
          if not root:
              return root
          root.left,root.right=root.right,root.left
          self.invertTree(root.left)
          self.invertTree(root.right)
          return root
      
          # shorter recur
          if not root:
              return root
          root.left,root.right=self.invertTree(root.right),self.invertTree(root.left)
          return root
  ```

- [200. Number of Islands](https://leetcode.cn/problems/number-of-islands/)

  ```python
  class Solution:
      def numIslands(self, grid: List[List[str]]) -> int:
          m,n=len(grid),len(grid[0])
          cnt=0
          def dfs(x,y):
              if x>=m or x<0 or y>=n or y<0:
                  return
              if grid[x][y]=='0' or grid[x][y]=='2':
                  return
              grid[x][y]='2'
              neighbour=[(0,1),(0,-1),(1,0),(-1,0)]
              for dx,dy in neighbour:
                  dfs(x+dx,y+dy)
          for i in range(m):
              for j in range(n):
                  if grid[i][j]!='0' and grid[i][j]!='2': # 2=searched
                      dfs(i,j)
                      cnt+=1
          return cnt   
  ```


## 20260410

- [207. Course Schedule](https://leetcode.cn/problems/course-schedule/)

  ```python
  class Solution:
      def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
          # no loop in graph
          # [0,1]
          # [1,0]
          # a->b->a
          # a->b->c->a
          # import numpy
          # relation=numpy.zeros((numCourses,numCourses))
          # relation=[[0]*numCourses for _ in range(numCourses)]
          # for ai, bi in prerequisites:
          #     # bi->ai
          #     relation[bi][ai]=1
  
          # Relation Matrix
          # temp=relation
          # for _ in range(numCourses):
          #     for i in range(numCourses):
          #         if temp[i][i]>0:
          #             return False
          #     temp=temp@relation
          # return True
  
          # Floyd
          # for k in range(numCourses):
          #     for i in range(numCourses):
          #         for j in range(numCourses):
          #             relation[i][j]=relation[i][j] or (relation[i][k] and relation [k][j])
          #     if relation[k][k]:
          #         return False
          # return True
  
          # Topo Sort
          from collections import deque
          in_degree=[0]*numCourses # if 0 then enque
          follower=[[] for _ in range(numCourses)]
          for ai,bi in prerequisites:
              in_degree[ai]+=1
              follower[bi].append(ai)
          cnt=0
          deq=deque([i for i in range(numCourses) if in_degree[i]==0])
          while deq:
              course=deq.popleft()
              cnt+=1
              if follower[course]:
                  for c in follower[course]:
                      in_degree[c]-=1
                      if in_degree[c]==0:
                          deq.append(c) # append apendleft pop popleft
          return cnt==numCourses  
  ```


## 20260411

- [300. Longest Increasing Subsequence](https://leetcode.cn/problems/longest-increasing-subsequence/)

  ```python
  class Solution:
      def lengthOfLIS(self, nums: List[int]) -> int:
  
          # dp
          '''
          if not nums:
              return 0
          n=len(nums)
          dp=[1]*n # max_len ends with idx i
          for i in range(n):
              max_pre=0
              for j in range(i): # find the best pre with largest dp val
                  if nums[i]>nums[j]:
                      max_pre=max(max_pre,dp[j])
              dp[i]=max_pre+1
          # return dp[n-1]
          return max(dp)
          '''
  
          # dp optimize
          if not nums:
              return 0
          n=len(nums)
          dp=[1]*n # max_len ends with idx i
          for i in range(n):
              for j in range(i): # find best pre with largest dp val
                  if nums[i]>nums[j]:
                      dp[i]=max(dp[i],dp[j]+1)
          return max(dp)
  
          # DP 核心问题：这个问题的上一个问题是什么
  ```


## 20260412

- [139. Word Break](https://leetcode.cn/problems/word-break/)

  ```python
  class Solution:
      def wordBreak(self, s: str, wordDict: List[str]) -> bool:
          if not s or not wordDict:
              return False
          n=len(s)
          wordDict=set(wordDict) # O(1) lookup optimize
          @cache
          def dp(i): # valid till i (0...n)
              if i==0:
                  return True
              for w in wordDict:
                  if i>=len(w) and w==s[i-len(w):i] and dp(i-len(w)):
                      return True
              return False
          return dp(n)
  ```


## 20260413

- [322. Coin Change](https://leetcode.cn/problems/coin-change/)

  ```python
  class Solution:
      def coinChange(self, coins: List[int], amount: int) -> int:
        if not coins or not amount:
          return 0
        @cache
        def dp(amt):
          if amt==0:
            return 0
          res=float('inf')
          for c in coins:
            if amt-c>=0 and dp(amt-c)!=-1:
              res=min(res,dp(amt-c)+1)
          return res if res!=float('inf') else -1
        return dp(amount)
  ```


## 20260414

- [17. Letter Combinations of a Phone Number](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/)

  ```python
  class Solution:
      def letterCombinations(self, digits: str) -> List[str]:
          if not digits:
              return []
          n=len(digits)
          mp={'2':['a','b','c'],'3':['d','e','f'],'4':['g','h','i'],'5':['j','k','l'],'6':['m','n','o'],'7':['p','q','r','s'],'8':['t','u','v'],'9':['w','x','y','z']}
  
          # it bfs
          '''
          res=['']
          for d in digits:
              cands=mp[d]
              res_new=[]
              for r in res:
                  for c in cands:
                      res_new.append(r+c)
              res=res_new
          return res
          '''
  
          # dp recur bfs
          '''
          def dp(idx,res):
              if idx==n:
                  return res
              cands=mp[digits[idx]]
              res_new=[]
              for r in res:
                  for c in cands:
                      res_new.append(r+c)
              return dp(idx+1,res_new)
          return dp(0,[''])
          '''
  
          # bt dfs
          def bt(idx):
              if idx==n:
                  rollouts.append(''.join(rollout))
                  return
              for c in mp[digits[idx]]:
                  rollout.append(c)
                  bt(idx+1)
                  rollout.pop()
          rollouts=[]
          rollout=[]
          bt(0)
          return rollouts
  ```


## 20260415

- [104. Maximum Depth of Binary Tree](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, val=0, left=None, right=None):
  #         self.val = val
  #         self.left = left
  #         self.right = right
  class Solution:
      def maxDepth(self, root: Optional[TreeNode]) -> int:
          # bfs
          if not root:
              return 0
          que=deque([root])
          cnt=0
          while que:
              curLev=len(que)
              cnt+=1
              for i in range(curLev):
                  node=que.popleft()
                  if node.left:
                      que.append(node.left)
                  if node.right:
                      que.append(node.right)
          return cnt
  ```

- [101. Symmetric Tree](https://leetcode.cn/problems/symmetric-tree/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, val=0, left=None, right=None):
  #         self.val = val
  #         self.left = left
  #         self.right = right
  class Solution:
      def isSymmetric(self, root: Optional[TreeNode]) -> bool:
          # bfs
          if not root:
              return False
          # if not root.left and not root.right:
          #     return True
          def mirror(a,b):
              if not a and not b:
                  return True
              if not a or not b:
                  return False
              return a.val==b.val and mirror(a.left,b.right) and mirror(a.right,b.left)
          # if root.left and root.right and root.left.val==root.right.val:
          #     return mirror(root.left,root.right)
          # return False
          return mirror(root.left,root.right)
  ```

- [108. Convert Sorted Array to Binary Search Tree](https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, val=0, left=None, right=None):
  #         self.val = val
  #         self.left = left
  #         self.right = right
  class Solution:
      def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
          if not nums:
              return None
          n=len(nums)
          mid=n//2
          root=TreeNode(val=nums[mid],left=self.sortedArrayToBST(nums[:mid]),right=self.sortedArrayToBST(nums[mid+1:]))
          return root
  ```


## 20260416

- [46. Permutations](https://leetcode.cn/problems/permutations/)

  ```python
  class Solution:
      def permute(self, nums: List[int]) -> List[List[int]]:
  
          # if not nums:
          #     return None
          # res=[]
          # rollout=[]
          # def bt(idx):
          #     if idx==len(nums):
          #         res.append(rollout[:])
          #         return
          #     for n in nums:
          #         if n not in rollout:
          #             rollout.append(n)
          #             bt(idx+1)
          #             rollout.pop()
          # bt(0)
          # return res
          
          # optimize lookup
          if not nums:
              return None
          res=[]
          rollout=[]
          visited=[False]*len(nums)
          def bt():
              if len(rollout)==len(nums):
                  res.append(rollout[:])
                  return
              for i in range(len(nums)):
                  if not visited[i]:
                      visited[i]=True
                      rollout.append(nums[i])
                      bt()
                      rollout.pop()
                      visited[i]=False
          bt()
          return res
  ```


## 20260417

- [39. Combination Sum](https://leetcode.cn/problems/combination-sum/)

  ```python
  class Solution:
      def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
          if not candidates or not target:
              return [[]]
          res=set()
          rollout=[]
          n=len(candidates)
          def bt():
              cur=sum(rollout)
              if cur==target:
                  res.add(tuple(sorted(rollout[:])))
                  return
              for i in range(n):
                  if cur+candidates[i]<=target:
                      rollout.append(candidates[i])
                      bt()
                      rollout.pop()
          bt()
          return list(res)
  ```


## 20260418

- [22. Generate Parentheses](https://leetcode.cn/problems/generate-parentheses/)

  ```python
  class Solution:
      def generateParenthesis(self, n: int) -> List[str]:
          if n==0:
              return [""]
          res=[]
          rollout=[]
          '''
          def bt(l=0,r=0):
              if l<r: # ensure l>=r
                  return
              if l==n and r==n:
                  res.append(''.join(rollout))
                  return
              if l<=n:
                  # add left
                  l+=1
                  rollout.append('(')
                  bt(l,r)
                  rollout.pop()
                  l-=1
                  # add right
                  r+=1
                  rollout.append(')')
                  bt(l,r)
                  rollout.pop()
                  r-=1
          '''
          # optimize
          def bt(l=0,r=0):
              if l<r or l>n: # not ((l>=r) and (l<=n))
                  return
              if l==n and r==n:
                  res.append(''.join(rollout))
                  return
              # add left
              rollout.append('(')
              bt(l+1,r)
              rollout.pop()
              # add right
              rollout.append(')')
              bt(l,r+1)
              rollout.pop()
          bt()
          return res
  ```


- [62. Unique Paths](https://leetcode.cn/problems/unique-paths/)

  ```python
  class Solution:
      def uniquePaths(self, m: int, n: int) -> int:
          @cache
          def dp(x,y):
              if x==0 or y==0:
                  return 1
              # if x>0 and y>0:
              #     return dp(x-1,y)+dp(x,y-1)
              # if x>0:
              #     return dp(x-1,y)
              # if y>0:
              #     return dp(x,y-1)
              return dp(x-1,y)+dp(x,y-1)
          return dp(m-1,n-1)
  ```


## 20260419

- [79. Word Search](https://leetcode.cn/problems/word-search/)

  ```python
  class Solution:
      '''
      def exist(self, board: List[List[str]], word: str) -> bool:
          if not board or not word:
              return False
          n=len(board)
          m=len(board[0])
          visited=[[False]*m for _ in range(n)]
          rollout=[]
          def bt(idx,ix,iy):
              if idx==len(word) and ''.join(rollout)==word:
                  return True
              if idx>=len(word):
                  return False
              for x,y in [(ix-1,iy),(ix+1,iy),(ix,iy+1),(ix,iy-1)]:
                  if x<0 or x>=n or y<0 or y>=m:
                      continue
                  if board[x][y]==word[idx] and not visited[x][y]:
                      visited[x][y]=True
                      rollout.append(word[idx])
                      if bt(idx+1,x,y):
                          return True
                      rollout.pop()
                      visited[x][y]=False
              return False
          for i in range(n):
              for j in range(m):
                  if board[i][j]==word[0]:
                      visited[i][j]=True
                      rollout.append(word[0])
                      if bt(1,i,j):
                          return True
                      rollout.pop()
                      visited[i][j]=False
          return False
      '''
  
      # optimize
      '''
      def exist(self, board: List[List[str]], word: str) -> bool:
          if not board or not word:
              return False
          n=len(board)
          m=len(board[0])
          visited=[[False]*m for _ in range(n)]
          def bt(idx,i,j):
              if idx==len(word):
                  return True
              for x,y in [(i-1,j),(i+1,j),(i,j+1),(i,j-1)]:
                  if x<0 or x>=n or y<0 or y>=m:
                      continue
                  if board[x][y]==word[idx] and not visited[x][y]:
                      visited[x][y]=True
                      if bt(idx+1,x,y):
                          return True
                      visited[x][y]=False
              return False
          for i in range(n):
              for j in range(m):
                  if board[i][j]==word[0]:
                      visited[i][j]=True
                      if bt(1,i,j):
                          return True
                      visited[i][j]=False
          return False
      '''
  
      # optimize
      def exist(self, board: List[List[str]], word: str) -> bool:
          if not board or not word:
              return False
          n=len(board)
          m=len(board[0])
          visited=[[False]*m for _ in range(n)]
          def bt(idx,i,j): # validator
              if idx==len(word):
                  return True
              if i<0 or i>=n or j<0 or j>=m:
                  return False
              if board[i][j]!=word[idx]:
                  return False
              if visited[i][j]:
                  return False
              visited[i][j]=True
              for x,y in [(i-1,j),(i+1,j),(i,j-1),(i,j+1)]:
                  if bt(idx+1,x,y):
                      return True
              visited[i][j]=False
              return False
          for i in range(n):
              for j in range(m):
                  if bt(0,i,j):
                      return True
                  # if board[i][j]==word[0]:
                  #     if bt(0,i,j):
                  #         return True
          return False
  ```


## 20260420

- [64. Minimum Path Sum](https://leetcode.cn/problems/minimum-path-sum/)

  ```python
  class Solution:
      def minPathSum(self, grid: List[List[int]]) -> int:
          '''
          n=len(grid)
          m=len(grid[0])
          @cache
          def dp(x,y):
              if x==0 and y==0:
                  return grid[0][0]
              if x>(n-1) or x<0 or y>(m-1) or y<0:
                  return float('inf')
              min_sum=float('inf')
              for pre_x,pre_y in [(x-1,y),(x,y-1)]:
                  min_sum=min(min_sum,grid[x][y]+dp(pre_x,pre_y))
              return min_sum
          return dp(n-1,m-1)
          '''
          # optimize
          n=len(grid)
          m=len(grid[0])
          @cache
          def dp(x,y):
              if x==0 and y==0:
                  return grid[0][0]
              # if x>(n-1) or x<0 or y>(m-1) or y<0:
              if x<0 or y<0:
                  return float('inf')
              # min_sum=float('inf')
              # for pre_x,pre_y in [(x-1,y),(x,y-1)]:
              #     min_sum=min(min_sum,grid[x][y]+dp(pre_x,pre_y))
              # return min_sum
              return grid[x][y]+min(dp(x-1,y),dp(x,y-1))
          return dp(n-1,m-1)
  ```


## 20260421

- [5. Longest Palindromic Substring](https://leetcode.cn/problems/longest-palindromic-substring/)

  ```python
  class Solution:
      def longestPalindrome(self, s: str) -> str:
          # bf
          '''
          n=len(s)
          if n==1:
              return s
          res=''
          for i in range(n):
              for j in range(i,n):
                  if j-i+1>len(res) and s[i:j+1]==s[i:j+1][::-1]:
                      res=s[i:j+1]
          return res
          '''
  
          # core expand
          def expand(l,r):
              while l>=0 and r<len(s) and s[l]==s[r]:
                  l-=1
                  r+=1
              # return s[l:r+1] broke
              return s[l+1:r]
  
          res=''
          for i in range(len(s)):
              single=expand(i,i)
              double=expand(i,i+1)
              res=max(res,single,double,key=len) # key=len
          return res
  ```


## 20260421

- [72. Edit Distance](https://leetcode.cn/problems/edit-distance/)

  ```python
  class Solution:
      def minDistance(self, word1: str, word2: str) -> int:
          '''
          @cache
          def dp(w1,w2):
              if w1==w2:
                  return 0
              if not w1:
                  return(len(w2))
              if not w2:
                  return(len(w1))
              min_cnt=float('inf')
              for i in range(len(w1)): # how to manipulate this single ori char
                  for j in range(len(w2)): # the target char
                      if w1[i]==w2[j]:
                          # remain
                          cnt_rem=dp(w1[:i],w2[:j])+dp(w1[i+1:],w2[j+1:])
                          min_cnt=min(min_cnt,cnt_rem)
                      else:
                          # replace
                          cnt_rep=1+dp(w1[:i],w2[:j])+dp(w1[i+1:],w2[j+1:])
                          # del
                          cnt_del_1=1+dp(w1[:i],w2[:j])+dp(w1[i+1:],w2[j:]) # i-1) j-1) (i+1 (j
                          cnt_del_2=1+dp(w1[:i],w2[:j+1])+dp(w1[i+1:],w2[j+1:]) # i-1) j) (i+1 (j+1
                          # add
                          cnt_add_1=1+dp(w1[:i+1],w2[:j])+dp(w1[i+1:],w2[j:]) # i) j-1) (i+1 (j
                          cnt_add_2=1+dp(w1[:i-2],w2[:j+1])+dp(w1[i:],w2[j+1:]) # i-1) j) (i (j+1
                          min_cnt=min(min_cnt,cnt_rep,cnt_del_1,cnt_del_2,cnt_add_1,cnt_add_2)
              return min_cnt
          return dp(word1,word2)
          '''
          @cache
          def dp(i,j): # up to
              if i==len(word1) and j==len(word2):
                  return 0
              if i==len(word1):
                  return len(word2)-j
              if j==len(word2):
                  return len(word1)-i
  
              if word1[i]==word2[j]:
                  return dp(i+1,j+1)
              else:
                  min_cnt=float('inf')
                  cnt_rep=1+dp(i+1,j+1)
                  cnt_del=1+dp(i+1,j)
                  cnt_add=1+dp(i,j+1)
                  min_cnt=min(min_cnt,cnt_rep,cnt_del,cnt_add)
              return min_cnt
          return dp(0,0) # forward
  ```

- [53. Maximum Subarray](https://leetcode.cn/problems/maximum-subarray/)

  ```python
  class Solution:
      def maxSubArray(self, nums: List[int]) -> int:
          max_sum=float('-inf')
          @cache
          def dp(idx):
              nonlocal max_sum
              if idx==0:
                  cur_sum=max_sum=nums[0]
                  return cur_sum
              cur_sum=max(nums[idx],nums[idx]+dp(idx-1))
              max_sum=max(max_sum,cur_sum)
              return cur_sum
          dp(len(nums)-1)
          return max_sum
  ```


## 20260423

- [283. Move Zeroes](https://leetcode.cn/problems/move-zeroes/)

  ```python
  class Solution:
      def moveZeroes(self, nums: List[int]) -> None:
          """
          Do not return anything, modify nums in-place instead.
          """
          # bf
          '''
          target=[n for n in nums if n!=0]+[n for n in nums if n==0]
          for i in range(len(nums)):
              nums[i]=target[i]  
          '''
  
          # swap
          '''
          if not nums or len(nums)==1:
              return
          non_0_idx=0
          for i in range(len(nums)):
              if nums[i]!=0:
                  nums[i],nums[non_0_idx]=nums[non_0_idx],nums[i]
                  non_0_idx+=1
          '''
  
          # prefill
          # filter 0
          non_0_idx=0
          for i in range(len(nums)):
              if nums[i]!=0:
                  nums[non_0_idx]=nums[i]
                  non_0_idx+=1
          # compensate 0
          for i in range(non_0_idx,len(nums)):
              nums[i]=0
  ```


## 20260424

- [35. Search Insert Position](https://leetcode.cn/problems/search-insert-position/)

  ```python
  class Solution:
      def searchInsert(self, nums: List[int], target: int) -> int:
          # lib
          # return bisect.bisect_left(nums,target)
          
          # manual
          less=[x for x in nums if x < target]
          return len(less)
  ```


## 20260425

- [20. Valid Parentheses](https://leetcode.cn/problems/valid-parentheses/)

  ```python
  class Solution:
      def isValid(self, s: str) -> bool:
          '''
          stk=[]
          for i in range(len(s)):
              if s[i]=='(' or s[i]=='[' or s[i]=='{':
                  stk.append(s[i])
              elif s[i]==')':
                  if not stk or stk[-1]!='(':
                      return False
                  stk.pop()
              elif s[i]==']':
                  if not stk or stk[-1]!='[':
                      return False
                  stk.pop()
              elif s[i]=='}':
                  if not stk or stk[-1]!='{':
                      return False
                  stk.pop()
          return not stk
          '''
          
          # optimize
          stk=[]
          pairs={'(':')','[':']','{':'}'}
          for c in s:
              # if c in pairs.keys():
              if c in pairs:
                  stk.append(c)
              elif not stk or c!=pairs[stk.pop()]:
                  return False
          return not stk
  ```


## 20260428

- [136. Single Number](https://leetcode.cn/problems/single-number/)

  ```python
  class Solution:
      def singleNumber(self, nums: List[int]) -> int:
          res=0
          for num in nums:
              res ^= num
          return res
  ```

- [169. Majority Element](https://leetcode.cn/problems/majority-element/)

  ```python
  class Solution:
      def majorityElement(self, nums: List[int]) -> int:
          # Set
          '''
          return max(set(nums),key=nums.count)
          '''
      	# Counter
          '''
      	from collections import Counter
      	return Counter(nums).most_common()[0][0] #[('apple', 3), ('banana', 2)]
      	'''
          # Voting
          cnt=0
          cand=None
          for num in nums:
              if cnt==0:
                  cand=num
              cnt+=(1 if cand==num else -1)
          return cand
      	
  ```


## 20260429

- [236. Lowest Common Ancestor of a Binary Tree](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, x):
  #         self.val = x
  #         self.left = None
  #         self.right = None
  
  class Solution:
      def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
          # recur
          '''
          if not root:
              return None # return nothing
          if root==p or root==q:
              return root # return the found one (directedly itself)
          l=self.lowestCommonAncestor(root.left,p,q) # leftsearch p | q
          r=self.lowestCommonAncestor(root.right,p,q) # rightsearch p | q
          if l and r: # could find from both side
              return root
          return l if l else r # return the one found (indirectedly)
          '''
  
          # iter bfs
          que=deque([root])
          parent={root:None}
          while que:
              n=que.popleft()
              if n.left:
                  parent[n.left]=n
                  que.append(n.left)
              if n.right:
                  parent[n.right]=n
                  que.append(n.right)
          ancestors=set()
          while p:
              ancestors.add(p)
              p=parent[p]
          while q not in ancestors:
              q=parent[q]
          return q
  ```


## 20260508

- [49. Group Anagrams](https://leetcode.cn/problems/group-anagrams/)

  ```python
  class Solution:
      def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
          # mapping
          n=len(strs)
          maps=defaultdict(list)
          for i in range(n):
              s=strs[i]
              # maps[tuple(sorted(s))].append(s)
              sid=[0]*26
              for c in s:
                  sid[ord(c)-ord('a')]+=1
              maps[tuple(sid)].append(s)
          return list(maps.values())
  ```


## 20260513

- [128. Longest Consecutive Sequence](https://leetcode.cn/problems/longest-consecutive-sequence/)

  ```python
  class Solution:
      def longestConsecutive(self, nums: List[int]) -> int:
          # bf
          '''
          maxseq=0
          curseq=0
          nums.sort()
          for i in range(len(nums)):
              if i==0 or nums[i]==nums[i-1]+1:
                  curseq+=1
              elif nums[i]==nums[i-1]:
                  pass
              else:
                  curseq=1 # reset
              maxseq=max(curseq,maxseq)
          return maxseq
          '''
          # unifind o(n)
          num_set=set(nums)
          max_len=0
          for num in num_set:
              if num-1 not in num_set: # min of the seq
                  cur_len=1
                  while num+1 in num_set:
                      num+=1
                      cur_len+=1
                  max_len=max(max_len,cur_len)
          return max_len
  ```


## 20260514

- [11. Container With Most Water](https://leetcode.cn/problems/container-with-most-water/)

  ```python
  class Solution:
      def maxArea(self, height: List[int]) -> int:
          start,end=0,len(height)-1 # 最悲观方案
          max_area=0
          while start < end: # 边界收缩
              area=min(height[end],height[start])*(end-start)
              if area > max_area:
                  max_area=area
              if height[start]<height[end]: # 优化短板
                  start+=1
              else:
                  end-=1
          return max_area
  ```

- [42. Trapping Rain Water](https://leetcode.cn/problems/trapping-rain-water/)

  ```python
  class Solution:
      def trap(self, height: List[int]) -> int:
          # bf
          '''
          max_height=max(height)
          max_width=len(height)
          res=0
          for h in range(max_height,0,-1):
              cand=[]
              for i in range(max_width):
                  if height[i]>=h:
                      cand.append(i)
              if len(cand)>=2:
                  for j in range(1,len(cand)):
                      dist=cand[j]-cand[j-1]-1
                      if dist>0:
                          res+=dist
          return res
          '''
          # calculus
          s,e=0,len(height)-1
          max_h_s=max_h_e=0
          res=0
          while s<e:
              if height[s]<height[e]:
                  if height[s]>max_h_s:
                      max_h_s=height[s]
                  else:
                      res+=(max_h_s-height[s])
                  s+=1
              else:
                  if height[e]>max_h_e:
                      max_h_e=height[e]
                  else:
                      res+=(max_h_e-height[e])
                  e-=1
          return res
  ```

## 20260515

- [54. Spiral Matrix](https://leetcode.cn/problems/spiral-matrix/)

  ```python
  class Solution:
      def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
          if not matrix:
              return matrix
          # collect
          res=[]
          while matrix:
              for i in matrix[0]:
                  res.append(i)
              # del
              matrix=matrix[1:]
              # left 90
              if matrix:
                  m=len(matrix)
                  n=len(matrix[0])
                  new_mat=[]
                  for j in range(n-1,-1,-1):
                      row=[]
                      for i in range(m):
                          row.append(matrix[i][j])
                      new_mat.append(row)
                  matrix=new_mat
          return res
  ```


## 20260519

- [56. Merge Intervals](https://leetcode.cn/problems/merge-intervals/)

  ```python
  class Solution:
      def merge(self, intervals: List[List[int]]) -> List[List[int]]:
          intervals.sort(key=lambda x:x[0])
          res=[intervals[0]]
          for i in range(1,len(intervals)):
              if intervals[i][1]>res[-1][1]:
                  if intervals[i][0]<=res[-1][1]:
                      res[-1][1]=intervals[i][1] # prolong
                  else:
                      res.append(intervals[i]) # new period
          return res
  ```


## 20260520

- [189. Rotate Array](https://leetcode.cn/problems/rotate-array/)

  ```python
  class Solution:
      def rotate(self, nums: List[int], k: int) -> None:
          """
          Do not return anything, modify nums in-place instead.
          """
          n=len(nums)
          k=k%n
          # temp=nums[n-k:]+nums[:n-k]
          # for i in range(n):
          #     nums[i]=temp[i]
          nums[:]=nums[-k:]+nums[:-k] # [:] in-place replace
  ```

- [238. Product of Array Except Self](https://leetcode.cn/problems/product-of-array-except-self/)

  ```python
  class Solution:
      def productExceptSelf(self, nums: List[int]) -> List[int]:
          res=[]
          n=len(nums)
          mul_forward=[1]*n
          mul_backward=[1]*n
          for i in range(1,n):
              mul_forward[i]=mul_forward[i-1]*nums[i-1]
              mul_backward[-i-1]=mul_backward[-i]*nums[-i]
          for i in range(n):
              # mul=1
              # for j in range(n):
              #     if j!=i:
              #         mul*=nums[j]
              # res.append(mul)
              res.append(mul_forward[i]*mul_backward[i])
          return res
  ```


## 20260521

- [73. Set Matrix Zeroes](https://leetcode.cn/problems/set-matrix-zeroes/)

  ```python
  class Solution:
      def setZeroes(self, matrix: List[List[int]]) -> None:
          """
          Do not return anything, modify matrix in-place instead.
          """
          rows=set()
          cols=set()
          for i in range(len(matrix)):
              for j in range(len(matrix[0])):
                  if matrix[i][j]==0:
                      rows.add(i)
                      cols.add(j)
          for i in rows:
              # for j in range(len(matrix[0])):
              #     matrix[i][j]=0
              matrix[i]=[0]*len(matrix[0])
          for i in range(len(matrix)):
              for j in cols:
                  matrix[i][j]=0
  ```

- [48. Rotate Image](https://leetcode.cn/problems/rotate-image/)

  ```python
  class Solution:
      def rotate(self, matrix: List[List[int]]) -> None:
          """
          Do not return anything, modify matrix in-place instead.
          """
          n=len(matrix)
          '''
          new_mat=[[0]*n for _ in range(n)]
          for i in range(n):
              for j in range(n):
                  new_mat[i][j]=matrix[-j-1][i]
          for i in range(n):
              for j in range(n):
                  matrix[i][j]=new_mat[i][j]
          '''
          # in-place replace
          # transpose
          for i in range(n):
              for j in range(i+1,n): # above diag
                  matrix[i][j],matrix[j][i]=matrix[j][i],matrix[i][j]
          # mirror left-right
          for i in range(n):
              for j in range(n//2): # left of centre vertical line
                  matrix[i][j],matrix[i][-j-1]=matrix[i][-j-1],matrix[i][j]
  ```


## 20260527

- [76. Minimum Window Substring](https://leetcode.cn/problems/minimum-window-substring/)

  ```python
  class Solution:
      def minWindow(self, s: str, t: str) -> str:
          # bf
          '''
          dict_t=defaultdict(int)
          for c in t:
              dict_t[c]+=1
          pre=pos=0
          dict_s=defaultdict(int)
          def check():
              nonlocal dict_s
              nonlocal dict_t
              for k in dict_t:
                  if dict_s[k]<dict_t[k]:
                      return False
              return True
          cand=[]
          while pre<len(s):
              if s[pre] in dict_t:
                  dict_s[s[pre]]+=1
                  if check():
                      # depulicate
                      while pos<=pre:
                          if s[pos] in dict_t:
                              if dict_s[s[pos]]>dict_t[s[pos]]:
                                  dict_s[s[pos]]-=1
                                  pos+=1
                              else:
                                  break
                          else:
                              dict_s[s[pos]]-=1
                              pos+=1
                      heapq.heappush(cand,(pre-pos+1,pos,pre))
              pre+=1
          if cand: 
              _,l,r=heapq.heappop(cand)
              return s[l:r+1]
          else:
              return ""
          '''
  
          # optimize by maintaining suff & min
          dict_t=defaultdict(int)
          for c in t:
              dict_t[c]+=1
          pre=pos=0
          dict_s=defaultdict(int)
          min_len=inf
          min_pos_pre=(0,-1)
          suff=0
          while pre<len(s):
              if s[pre] in dict_t:
                  dict_s[s[pre]]+=1
                  if dict_s[s[pre]]==dict_t[s[pre]]:
                      suff+=1
                      # depulicate
                      while pos<=pre and suff==len(dict_t):
                          cur_len=pre-pos+1
                          if cur_len<min_len:
                              min_len=cur_len
                              min_pos_pre=(pos,pre)
                          if s[pos] in dict_t:
                              dict_s[s[pos]]-=1
                              if dict_s[s[pos]]<dict_t[s[pos]]:
                                  suff-=1
                          pos+=1
              pre+=1
          return s[min_pos_pre[0]:min_pos_pre[1]+1]
  ```


## 20260528

- [155. Min Stack](https://leetcode.cn/problems/min-stack/)

  ```python
  class MinStack:
  
      # hq
      '''
      def __init__(self):
          self.stk=[]
          self.hq=[]
          # heapq.heapify(self.hq)
  
      def push(self, val: int) -> None:
          self.stk.append(val)
          heapq.heappush(self.hq,val)
          
      def pop(self) -> None:
          el=self.stk[-1]
          self.hq.remove(el)
          heapq.heapify(self.hq) # re-adjust the tree
          self.stk.pop()
          
      def top(self) -> int:
          return self.stk[-1]
  
      def getMin(self) -> int:    
          return self.hq[0]
      '''
  
      # minstk 同步各层最小值
      def __init__(self):
          self.stk=[]
          self.minstk=[]
  
      def push(self, val: int) -> None:
          self.stk.append(val)
          if not self.minstk or val<=self.minstk[-1]: # <= !!!
              self.minstk.append(val)
          
      def pop(self) -> None:
          if self.stk.pop()==self.minstk[-1]:
              self.minstk.pop()
          
      def top(self) -> int:
          return self.stk[-1]
  
      def getMin(self) -> int:    
          return self.minstk[-1]
  
  # Your MinStack object will be instantiated and called as such:
  # obj = MinStack()
  # obj.push(val)
  # obj.pop()
  # param_3 = obj.top()
  # param_4 = obj.getMin()
  ```


## 20260610

- [215. Kth Largest Element in an Array](https://leetcode.cn/problems/kth-largest-element-in-an-array/)

  ```python
  class Solution:
      def findKthLargest(self, nums: List[int], k: int) -> int:
          return heapq.nlargest(k,nums)[k-1]
  ```

- [295. Find Median from Data Stream](https://leetcode.cn/problems/find-median-from-data-stream/)

  ```python
  class MedianFinder:
  
      def __init__(self):
          self.xgd=[] # the greater half
          self.dgd=[] # the smaller half
  
      def addNum(self, num: int) -> None:
          heapq.heappush(self.dgd,-num)
          heapq.heappush(self.xgd,-heapq.heappop(self.dgd))
          if len(self.xgd)>len(self.dgd):
              # return back, ensuring dgd's size is larger
              heapq.heappush(self.dgd,-heapq.heappop(self.xgd))
  
      def findMedian(self) -> float:
          if len(self.dgd)>len(self.xgd):
              return -self.dgd[0]
          else:
              return (-self.dgd[0]+self.xgd[0])/2
  
  # Your MedianFinder object will be instantiated and called as such:
  # obj = MedianFinder()
  # obj.addNum(num)
  # param_2 = obj.findMedian()
  ```


## 20260611

- [124. Binary Tree Maximum Path Sum](https://leetcode.cn/problems/binary-tree-maximum-path-sum/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, val=0, left=None, right=None):
  #         self.val = val
  #         self.left = left
  #         self.right = right
  class Solution:
      def maxPathSum(self, root: Optional[TreeNode]) -> int:
          # kadane
          res=-inf
          def dfs(node):
              if not node:
                  return 0
              left=dfs(node.left)
              right=dfs(node.right)
              cur_max=node.val+max(0,left,right,left+right)
              # print(cur_max)
              nonlocal res
              res=max(res,cur_max)
              return node.val+max(0,left,right)
          part_max=dfs(root)
          res=max(res,part_max)
          return res
  
  #------------------------------------------------------#
  # standard
  class Solution:
      def maxPathSum(self, root: Optional[TreeNode]) -> int:
          res = -inf
  
          def dfs(node):
              nonlocal res
              if not node:
                  return 0
              
              # Kadane 思想：如果子树路径和是负数，我们直接放弃它（取 0）
              left_gain = max(dfs(node.left), 0)
              right_gain = max(dfs(node.right), 0)
              
              # 当前节点作为“最高拐点”（分叉点）时的最大路径和
              current_path_sum = node.val + left_gain + right_gain
              
              # 更新全局最大值
              res = max(res, current_path_sum)
              
              # 返回给父节点：当前节点能提供的最大“单侧”路径和
              return node.val + max(left_gain, right_gain)
  
          dfs(root)
          return res
  ```

- [146. LRU Cache](https://leetcode.cn/problems/lru-cache/)

  ```python
  class LRUCache:
  
      def __init__(self, capacity: int):
          self.capacity=capacity
          self.cnt=0
          self.kvcache=OrderedDict()
  
      def get(self, key: int) -> int:
          if key in self.kvcache:
              value=self.kvcache[key]
              self.kvcache.move_to_end(key)
              return value
          else:
              return -1
          
      def put(self, key: int, value: int) -> None:
          if key in self.kvcache:
              self.kvcache[key]=value
              self.kvcache.move_to_end(key)
          else:
              if self.cnt>=self.capacity:
                  # evict
                  self.kvcache.popitem(last=False)
                  self.cnt-=1
              # add
              self.kvcache[key]=value
              self.cnt+=1
  
  #-----------------------------------------------#
  # standard
   class _Node:
      """双向链表节点"""
      __slots__ = ('key', 'value', 'prev', 'next')
      def __init__(self, key=None, value=None):
          self.key = key
          self.value = value
          self.prev = None
          self.next = None
  
  class LRUCache:
      def __init__(self, capacity: int):
          if capacity <= 0:
              raise ValueError("capacity must be positive")
          self.capacity = capacity
          self.cache = {}            # key -> node
          # 哨兵节点，简化边界处理
          self.head = _Node()        # 虚拟头节点
          self.tail = _Node()        # 虚拟尾节点
          self.head.next = self.tail
          self.tail.prev = self.head
  
      def _remove_node(self, node: _Node) -> None: # del
          """从链表中移除一个节点（不删除缓存映射）"""
          prev_node = node.prev
          next_node = node.next
          prev_node.next = next_node
          next_node.prev = prev_node
  
      def _add_to_head(self, node: _Node) -> None: # add_h
          """将节点插入到头部（head 之后）"""
          node.prev = self.head
          node.next = self.head.next
          self.head.next.prev = node
          self.head.next = node
  
      def _move_to_head(self, node: _Node) -> None: # mv_h
          """将节点移动到头部（先移除再添加）"""
          self._remove_node(node)
          self._add_to_head(node)
  
      def _pop_tail(self) -> _Node: # pop_t
          """弹出尾部节点（即最后一个实际节点）"""
          node = self.tail.prev
          self._remove_node(node)
          return node
  
      def get(self, key: int) -> int:
          if key not in self.cache:
              return -1
          node = self.cache[key]
          # 访问后移到头部
          self._move_to_head(node)
          return node.value
  
      def put(self, key: int, value: int) -> None:
          if key in self.cache:
              # 更新已有节点的值并移到头部
              node = self.cache[key]
              node.value = value
              self._move_to_head(node)
          else:
              # 新节点
              if len(self.cache) >= self.capacity:
                  # 淘汰最久未使用的节点（尾部）
                  tail_node = self._pop_tail()
                  del self.cache[tail_node.key]
              # 创建新节点并插入头部
              new_node = _Node(key, value)
              self.cache[key] = new_node
              self._add_to_head(new_node)
  
  # Your LRUCache object will be instantiated and called as such:
  # obj = LRUCache(capacity)
  # param_1 = obj.get(key)
  # obj.put(key,value)
  ```


## 20260612

- [74. Search a 2D Matrix](https://leetcode.cn/problems/search-a-2d-matrix/)

  ```
  class Solution:
      def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
          # 1 1 2
          # 3 3 5
          m=len(matrix)
          n=len(matrix[0])
          x=0
          y=n-1
          while x<m and y>=0:
              if matrix[x][y]==target:
                  return True
              elif matrix[x][y]>target:
                  y-=1
              else:
                  x+=1
          return False
  ```

- [34. Find First and Last Position of Element in Sorted Array](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)

  ```python
  class Solution:
      def searchRange(self, nums: List[int], target: int) -> List[int]:
          # expand bound
          '''
          l=0
          r=len(nums)-1
          m=-1
          while l<=r:
              piv=l+(r-l)//2
              if nums[piv]==target:
                  m=piv
                  break
              elif nums[piv]>target:
                  r=piv-1
              else:
                  l=piv+1
          if m==-1:
              return [-1,-1]
          left=right=m
          while left-1>=0 and nums[left-1]==target:
              left-=1
          while right+1<=len(nums)-1 and nums[right+1]==target:
              right+=1
          return [left,right]
          '''
          # two pass
          def bsearch(leftest: bool) -> int:
              l=0
              r=len(nums)-1
              bound=-1
              while l<=r:
                  m=l+(r-l)//2
                  if nums[m]==target:
                      bound=m
                      if leftest:
                          r=m-1
                      else:
                          l=m+1
                  elif nums[m]>target:
                      r=m-1
                  else:
                      l=m+1
              return bound
          left=bsearch(leftest=True)
          right=bsearch(leftest=False)
          return [left,right]
  ```


## 20260613

- [98. Validate Binary Search Tree](https://leetcode.cn/problems/validate-binary-search-tree/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, val=0, left=None, right=None):
  #         self.val = val
  #         self.left = left
  #         self.right = right
  class Solution:
      def isValidBST(self, root: Optional[TreeNode]) -> bool:
          arr=[]
          def inorder(node: TreeNode) -> None:
              if not node:
                  return
              inorder(node.left)
              # nonlocal arr
              arr.append(node.val)
              inorder(node.right)
          inorder(root)
          if len(arr)<=1:
              return True
          # for i in range(len(arr)-1):
          #     if arr[i]>=arr[i+1]:
          #         return False
          # return True
          return all(arr[i]<arr[i+1] for i in range(len(arr)-1))
  ```

- [230. Kth Smallest Element in a BST](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, val=0, left=None, right=None):
  #         self.val = val
  #         self.left = left
  #         self.right = right
  class Solution:
      def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
          ctr=0
          res=-1
          def inorder(node: TreeNode) -> bool: # 表示找到没找到 提前终止以剪枝
              nonlocal ctr,res
              if not node:
                  return False
              if inorder(node.left):
                  return True
              ctr+=1
              if ctr==k:
                  res=node.val
                  return True
              return inorder(node.right)
          inorder(root)
  ```


## 20260614

- [199. Binary Tree Right Side View](https://leetcode.cn/problems/binary-tree-right-side-view/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, val=0, left=None, right=None):
  #         self.val = val
  #         self.left = left
  #         self.right = right
  class Solution:
      def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
          if not root:
              return []
          q=deque([root])
          res=[]
          while q:
              levlen=len(q)
              for i in range(levlen):
                  node=q.popleft()
                  if i==levlen-1:
                      res.append(node.val)
                  if node.left:
                      q.append(node.left)
                  if node.right:
                      q.append(node.right)
          return res
  ```

- [114. Flatten Binary Tree to Linked List](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, val=0, left=None, right=None):
  #         self.val = val
  #         self.left = left
  #         self.right = right
  class Solution:
      def flatten(self, root: Optional[TreeNode]) -> None:
          """
          Do not return anything, modify root in-place instead.
          """
  
          def flat(node: Optional[TreeNode]) -> Optional[TreeNode]:
              if not node:
                  return None
              if not node.left and not node.right:
                  return node
              '''
              left=flat(node.left)
              right=flat(node.right)
              if left:
                  left_last=left
                  while left_last.right:
                      left_last=left_last.right
                  left_last.right=right
                  node.left=None
                  node.right=left
              else:
                  node.left=None
                  node.right=right
  
              return node
              '''
              left_last=flat(node.left)
              right_last=flat(node.right)
              if left_last:
                  left_last.right=node.right
                  node.right=node.left
                  node.left=None
              return right_last if right_last else left_last
          
          flat(root)
  ```


## 20260615

- [105. Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

  ```python
  # Definition for a binary tree node.
  # class TreeNode:
  #     def __init__(self, val=0, left=None, right=None):
  #         self.val = val
  #         self.left = left
  #         self.right = right
  class Solution:
      def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
          if not preorder or not inorder or len(preorder)!=len(inorder):
              return None
          if len(preorder)==1:
              return TreeNode(val=preorder[0])
          root_val=preorder[0]
          root_idx_in=inorder.index(root_val)
          left_in=inorder[:root_idx_in]
          right_in=inorder[root_idx_in+1:]
          # left_pre=[v for v in preorder if v in left_in]
          # right_pre=[v for v in preorder if v in right_in]
          left_num=len(left_in)
          left_pre=preorder[1:1+left_num]
          right_pre=preorder[1+left_num:]
          left_node=self.buildTree(preorder=left_pre,inorder=left_in)
          right_node=self.buildTree(preorder=right_pre,inorder=right_in)
          root_node=TreeNode(val=root_val,left=left_node,right=right_node)
          return root_node
  ```

- [33. Search in Rotated Sorted Array](https://leetcode.cn/problems/search-in-rotated-sorted-array/)

  ```python
  class Solution:
      def search(self, nums: List[int], target: int) -> int:
  
          # bf
          '''
          # 确定旋转点
          n=len(nums)
          smallest_idx=0
          for i in range(n-1):
              if nums[i]>nums[i+1]:
                  smallest_idx=i+1
                  break
          
          # 确定区间
          if nums[smallest_idx]<=target and nums[-1]>=target:
              l=smallest_idx
              r=n-1
          elif nums[0]<=target and nums[smallest_idx-1]>=target:
              l=0
              r=(smallest_idx-1+n)%n
          else:
              return -1
  
          # 二分查找
          while l<=r:
              m=l+(r-l)//2
              if target==nums[m]:
                  return m
              elif target<nums[m]:
                  r=m-1
              else:   
                  l=m+1
          return -1
          '''
  
          # 有序区间+元素二分查找
          l=0
          r=len(nums)-1
          while l<=r:
              m=l+(r-l)//2
              if target==nums[m]:
                  return m
              # 左有序
              elif nums[l]<=nums[m]: # 可能 l==m
                  # 又在左区间
                  if nums[l]<=target<nums[m]:
                      # 太好了 在左半段查找
                      r=m-1
                  else:
                      l=m+1
              # 右有序
              else:
                  # 又在右区间
                  if nums[m]<target<=nums[r]:
                      # 太好了 在右半段查找
                      l=m+1
                  else:
                      r=m-1
          return -1
  ```

- [153. Find Minimum in Rotated Sorted Array](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/)

  ```python
  class Solution:
      def findMin(self, nums: List[int]) -> int:
          # return min(nums)
          # bf
          '''
          for i in range(len(nums)-1):
              if nums[i]>nums[i+1]:
                  return nums[i+1]
          return nums[0]
          '''
          l=0
          r=len(nums)-1
          minv=inf
          while l<=r:
              m=l+(r-l)//2
              # 左有序
              if nums[l]<=nums[m]:
                  minv=min(minv,nums[l])
                  l=m+1
              # 右有序
              else:
                  minv=min(minv,nums[m])
                  r=m-1
          return minv
  # standard
  class Solution:
      def findMin(self, nums: List[int]) -> int:
          l, r = 0, len(nums) - 1
          while l < r:  # 注意这里是 l < r
              m = l + (r - l) // 2
              if nums[m] > nums[r]:
                  l = m + 1  # 最小值在右边
              else:
                  r = m      # m 有可能是最小值，所以不能 m-1
          return nums[l]
  ```

- [4. Median of Two Sorted Arrays](https://leetcode.cn/problems/median-of-two-sorted-arrays/)

  ```python
  class Solution:
      def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
          merged=[]
          p1=p2=0
          while p1<len(nums1) and p2<len(nums2):
              if nums1[p1]<nums2[p2]:
                  merged.append(nums1[p1])
                  p1+=1
              else:
                  merged.append(nums2[p2])
                  p2+=1
          if p1>=len(nums1):
              merged.extend(nums2[p2:]) 
          else:
              merged.extend(nums1[p1:])
          
          n=len(merged)
          if n%2==1:
              return merged[n//2]
          else:
              return (merged[n//2]+merged[n//2-1])/2
  ```

  
