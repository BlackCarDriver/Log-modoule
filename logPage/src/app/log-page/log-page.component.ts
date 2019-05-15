import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-log-page',
  templateUrl: './log-page.component.html',
  styleUrls: ['./log-page.component.css'],
  styles: []
})

export class LogPageComponent implements OnInit {
  private addr: string  = "http://localhost:8090";
  // private addr:string= "api";

  constructor(private http: HttpClient) { }
  ngOnInit() {}

  pageIndex = 1;  //操作日志首页码
  pageSize = 10;   //操作日志每页条数
  oreloglist = []   //操作日志
  
  // tslint:disable-next-line:member-ordering
  public orelogs: Orelog[];

  public orelog: Orelog = {
    logid: 0,
    logtime: '',
    admin:'',
    module:'',
    logsql:''
  };
  listOfData = [
    {
      key: '1',
      name: 'John Brown',
      age: 32,
      address: 'New York No. 1 Lake Park'
    },
    {
      key: '2',
      name: 'Jim Green',
      age: 42,
      address: 'London No. 1 Lake Park'
    },
    {
      key: '3',
      name: 'Joe Black',
      age: 32,
      address: 'Sidney No. 1 Lake Park'
    }
  ];
  // 请求操作日志内容
  public Getorelog() {           // 获取信息
    let url = this.addr + '/log/logdisplay';
    this.http.get<Orelog[]>(url).subscribe(response => {
      this.orelogs = response;
      //var testtrimresult = orelog.substring(0, testtrim.length-2);
    });
  }

  // 操作日志分页
  searchData(): void {
    this.oreloglist = this.orelogs.slice((this.pageIndex - 1) * this.pageSize, (this.pageIndex) * this.pageSize);
  }

}

class Orelog {
  public logid:       number;
  public logtime:     string;
  public admin:       string;
  public module:      string;
  public logsql:      string;
}
