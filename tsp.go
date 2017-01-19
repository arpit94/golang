package main

import "fmt"
const N = 10
var n int
var mygraph [N][N]int
var path [N+1]int
var result int

func main(){
  //mygraph = [N][N]int {{0,10,15,20},{10,0,35,25},{15,35,0,30},{20,25,30,0}}
  fmt.Print("Enter the number of vertices : ")
  fmt.Scan(&n)
  for p:=0;p<n;p++{
    for q:=0;q<n;q++{
      if p==q{
        mygraph[p][q]=0
      }else if p<q{
        fmt.Print("Enter the ",p,q," path cost : ")
        fmt.Scan(&mygraph[p][q])
        mygraph[q][p]=mygraph[p][q]
      }
    }
  }
  result=100000
  var visited [N]bool
  var currPath [N]int
  currBound :=0
  for i:=0;i<n;i++{
    currBound+=firstMin(i)+secondMin(i)
    visited[i]=false
    currPath[i]=-1
  }
  currPath[0]=0
  visited[0]=true
  TSP(currBound/2,currPath,0,1,visited)
  fmt.Println("Minimum cost of path  is : ",result)
  fmt.Print("Result path is : ",path[0])
  for i:=1;i<n+1;i++{
  	fmt.Print(",",path[i])
  }
}
func firstMin(r int) int{
  var min int
  if(r==0){
    min=mygraph[0][1]
  }else{
    min=mygraph[r][0]
  }
  for i:=1;i<n;i++{
    if (mygraph[r][i]<min) && (i!=r){
      min=mygraph[r][i]
    }
  }
  return min
}
func secondMin(r int) int{
  min1:=-1
  min2:=-1
  for j:=0;j<n;j++{
    if(r==j){
      //skip
    }else{
      if(min1==-1){
        min1=mygraph[r][j]
      }else{
        if(mygraph[r][j]<min1){
          min2=min1
          min1=mygraph[r][j]
        }else{
          if(min2==-1){
            min2=mygraph[r][j]
          }else{
            if(mygraph[r][j]<min2){
              min2=mygraph[r][j]
            }
          }
        }
      }
    }
  }
  return min2
}
func storePath(currPath [N]int){
  for i :=0;i<n;i++{
    path[i]=currPath[i]
    }
    path[n]=currPath[0]
}
func TSP(currBound int,currPath [N]int,currWeight int,level int,visited [N]bool){
  var tempRes int
  var temp int
  if (level==n)&&(mygraph[currPath[level-1]][currPath[0]]!=0){
	tempRes=currWeight+mygraph[currPath[level-1]][currPath[0]]
    if tempRes<result{
      result=tempRes
      storePath(currPath)
    }
  }else{
    for i:=0;i<n;i++{
      if (mygraph[currPath[level-1]][i] != 0) &&(visited[i] == false) {
        temp=currBound
        currWeight+=mygraph[currPath[level-1]][i]
        if level==1{
          currBound -= ((firstMin(currPath[level-1])+firstMin(i))/2)
        }else{
          currBound -= ((secondMin(currPath[level-1]) +firstMin(i))/2)
        }
        if (currBound + currWeight) < result{
          currPath[level] = i
          visited[i] = true
          TSP(currBound,currPath, currWeight, level+1,visited)
        }
        currWeight -= mygraph[currPath[level-1]][i]
        currBound = temp;
        visited[i]=false
      }
    }
  }
}
