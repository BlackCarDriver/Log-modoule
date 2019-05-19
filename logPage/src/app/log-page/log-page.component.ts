import { Component, OnInit, Pipe } from '@angular/core';
import { HttpClient} from '@angular/common/http';

@Component({
  selector: 'app-log-page',
  templateUrl: './log-page.component.html',
  styleUrls: ['./log-page.component.css'],
  styles: []
})

export class LogPageComponent implements OnInit {

  private addr: string  = "http://localhost:8090";

  totalrows:number=0;     //所需的类型共有多少行数据
  packindex : number=1;    //页码组件需要用到的当前页数值
  tpackindex :number=1;   //列表组件需要用到的当前页数值
  packlap : number = 0;   //页码数值与正在显示的数据所在的页数值差
  selecttype:string="all"; //加载日志的类型
  listOfData : log[] = [];      //日志
  pagerang : number[] = [0,0];       //已经加载的数据页数范围

  constructor(private http: HttpClient) { }

  ngOnInit() {
    this.getdata("all",this.packindex);
  }
  //获取日志列表，参数是类型和开始获取的页数
  //注意每页能够显示11行,每次从index开始请求10页数据，
  getdata(params:string, index:number){
      let url = this.addr + `/log/getlog?type=${params}&&index=${(index-1)*11}`;
      this.http.get<any>(url).subscribe(result=>{
       this.listOfData = result.log;
       this.totalrows = result.rowsnumber;
       this.packlap = index -1;
       this.pagerang[0] = index;
       this.pagerang[1] = index + parseInt(this.listOfData.length/11 + "");
      });
  }

  //由页数选择改变触发
  //每次缓存10页,范围是packrange,超过这个范围则再次请求数据
  updatapage(){
    if(this.packindex>=this.pagerang[0] && this.packindex < this.pagerang[1]){
      this.tpackindex = this.packindex - this.packlap;
      return;
    }
    this.getdata(this.selecttype, this.packindex);
  }
  //又类型选择改变触发
  updatatype(t:any){
    this.selecttype = t;
    this.getdata(t, 1);
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