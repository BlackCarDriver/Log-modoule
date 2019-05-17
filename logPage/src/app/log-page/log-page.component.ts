import { Component, OnInit } from '@angular/core';
import { HttpClient} from '@angular/common/http';

@Component({
  selector: 'app-log-page',
  templateUrl: './log-page.component.html',
  styleUrls: ['./log-page.component.css'],
  styles: []
})

export class LogPageComponent implements OnInit {

  private addr: string  = "http://localhost:8090";

  totalrowsr:number=0;     //数据总数
  haverows:number=0;       //当前已有数据行数
  rowsnumber :number = 0;  //返回的数据一共有多少行
  packindex : number=1;    //当前的所在页数
  selecttype:string="all"; //加载日志的类型
  listOfData : log[] = [];      //日志

  constructor(private http: HttpClient) { }

  ngOnInit() {
    this.getdata("warn",this.haverows);
  }
  //获取日志列表，参数是类型和开始获取的下标index，
  //注意每页能够显示11行,每次从index开始请求5页数据，
  getdata(params:string, index:number){
      let url = this.addr + `/log/getlog?type=${params}&&index=${index}`;
      this.http.get<any>(url).subscribe(result=>{
        console.log(result);
       this.listOfData = result.log;
       this.rowsnumber = result.rowsnumber;
      });
  }

  //由页数选择改变触发
  updatapage(){
    this.getdata(this.selecttype, 0);
  }
  //又类型选择改变触发
  updatatype(t:any){
    this.selecttype = t;
    this.getdata(t, 0);
  }

}

//日志元组结构体
class log{
   index:number;
   type:string;
   operator:string;
   time:string;
   operation:string;
}

//模拟数据
var listOfData = [
  {index:1, type: 'John Brown',  operator: "blackcardriver", time: '2019-11-22', operation:"登录系统" },
  {index:2, type: 'John Brown',  operator: "blackcardriver", time: '2019-11-22', operation:"登录系统" },
];