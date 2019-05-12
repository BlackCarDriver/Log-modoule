import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-log-page',
  templateUrl: './log-page.component.html',
  styleUrls: ['./log-page.component.css']
})
export class LogPageComponent implements OnInit {
  constructor() { }
    ngOnInit() {
  }

  panels = [
    { name: '2019-5-11',
      list:['第一次构建','第二次构建','第三次构建',]
  },
    { name: '2019-5-12',
    list:['第一次构建','第二次构建','第三次构建',]
},
{ name: '2019-5-11',
list:['第一次构建','第二次构建','第三次构建',]
},
{ name: '2019-5-12',
list:['第一次构建','第二次构建','第三次构建',]
},
{ name: '2019-5-11',
list:['第一次构建','第二次构建','第三次构建',]
},
{ name: '2019-5-12',
list:['第一次构建','第二次构建','第三次构建',]
}

  ];

}
