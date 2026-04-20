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

  
