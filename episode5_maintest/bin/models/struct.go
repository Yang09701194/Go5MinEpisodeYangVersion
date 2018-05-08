package models

import (
	"time"
)

type Seatconsolidationskills struct {
	Seatconsolidationruleid int64  `xorm:"not null pk BIGINT(8)"`
	Skillid                 []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Allowancemapping struct {
	Id       int    `xorm:"not null pk INT(4)"`
	Termname string `xorm:"not null NVARCHAR(100)"`
	Money    int    `xorm:"not null INT(4)"`
}

type Agentacdidcolor struct {
	Agentacdid string `xorm:"not null NVARCHAR(510)"`
	Color      string `xorm:"NVARCHAR(100)"`
}

type Sysdiagrams struct {
	Name        string `xorm:"not null unique(UK_principal_name) SYSNAME(256)"`
	PrincipalId int    `xorm:"not null unique(UK_principal_name) INT(4)"`
	DiagramId   int    `xorm:"not null pk INT(4)"`
	Version     int    `xorm:"INT(4)"`
	Definition  []byte `xorm:"VARBINARY(-1)"`
}

type Swapoffdute struct {
	Swapoffduteid    []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Appierid         []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Applydate        time.Time `xorm:"DATETIME(8)"`
	Applyoffdutetype string    `xorm:"NVARCHAR(100)"`
	Staus            string    `xorm:"NVARCHAR(32)"`
	Swapappdate      time.Time `xorm:"DATETIME(8)"`
}

type Biddingannouncementbiddinggroup struct {
	Biddingannouncementid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Biddinggroupid        int    `xorm:"not null pk INT(4)"`
	Biddinggroupname      string `xorm:"NVARCHAR(256)"`
}

type Swappost struct {
	Swappostid      []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Applicationdate time.Time `xorm:"not null DATETIME(8)"`
	Employeeid      []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Datea           time.Time `xorm:"not null DATETIME(8)"`
	Dateb           time.Time `xorm:"not null DATETIME(8)"`
	Assignmentid    int64     `xorm:"BIGINT(8)"`
	Assignmentinfo  string    `xorm:"NVARCHAR(510)"`
	Description     string    `xorm:"NVARCHAR(510)"`
	Status          string    `xorm:"not null NVARCHAR(32)"`
	Swaptype        string    `xorm:"not null NVARCHAR(32)"`
	Maturitydate    time.Time `xorm:"not null DATETIME(8)"`
	Iscutshift      int       `xorm:"default 0 BIT(1)"`
	Version         time.Time `xorm:"TIMESTAMP(8)"`
}

type Agentstatusalerttime struct {
	Agentstatusalerttimeid int    `xorm:"not null pk INT(4)"`
	Alerttimeoutsecond     int    `xorm:"not null INT(4)"`
	Alerttimeoutsecond2    int    `xorm:"not null INT(4)"`
	Agentstatustypecode    string `xorm:"NVARCHAR(510)"`
	Areaid                 []byte `xorm:"not null UNIQUEIDENTIFIER(16)"`
}

type Version struct {
	Version    string    `xorm:"CHAR(10)"`
	Updatedate time.Time `xorm:"DATETIME(8)"`
}

type Seatconsolidationorganizations struct {
	Seatconsolidationruleid int64  `xorm:"not null pk BIGINT(8)"`
	Organizationid          []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Biddingdistributionmode struct {
	Biddinggroupid int `xorm:"not null pk INT(4)"`
	Datepriority   int `xorm:"not null INT(4)"`
	Rank           int `xorm:"not null INT(4)"`
	Years          int `xorm:"not null INT(4)"`
	Iscustomized   int `xorm:"not null BIT(1)"`
}

type Term struct {
	Termid           int64     `xorm:"not null pk BIGINT(8)"`
	Childclasstype   string    `xorm:"not null NVARCHAR(64)"`
	Starttime        time.Time `xorm:"not null index(IX_TimeBox) DATETIME(8)"`
	Finishtime       time.Time `xorm:"not null index(IX_TimeBox) DATETIME(8)"`
	Locked           int       `xorm:"not null BIT(1)"`
	Levelcolumn      int       `xorm:"INT(4)"`
	Parenttermid     int64     `xorm:"BIGINT(8)"`
	Termstyleid      []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Starttime2       time.Time `xorm:"DATETIME(8)"`
	Finishtime2      time.Time `xorm:"DATETIME(8)"`
	Shrinkagetotals  int64     `xorm:"BIGINT(8)"`
	Overtimetotals   int64     `xorm:"BIGINT(8)"`
	Workingtotals    int64     `xorm:"BIGINT(8)"`
	Otstarttime2     time.Time `xorm:"DATETIME(8)"`
	Otfinishtime2    time.Time `xorm:"DATETIME(8)"`
	Otovertimetotals int64     `xorm:"BIGINT(8)"`
	Timeboxid        []byte    `xorm:"index(IX_TimeBox) UNIQUEIDENTIFIER(16)"`
	Tag              string    `xorm:"NVARCHAR(510)"`
	Version          int64     `xorm:"not null BIGINT(8)"`
	Seatid           string    `xorm:"NVARCHAR(510)"`
	Asarest          int       `xorm:"BIT(1)"`
	Asawork          int       `xorm:"BIT(1)"`
	Background       string    `xorm:"NVARCHAR(510)"`
	Belongtoprv      int       `xorm:"BIT(1)"`
	Displaytext      string    `xorm:"NVARCHAR(510)"`
	Gapguaranteed    int       `xorm:"BIT(1)"`
	Ignoreadherence  int       `xorm:"BIT(1)"`
	Isneedseat       int       `xorm:"BIT(1)"`
	Onservice        int       `xorm:"BIT(1)"`
	Custompayment    string    `xorm:"NVARCHAR(510)"`
	Nativename       string    `xorm:"NVARCHAR(510)"`
	Hrdate           time.Time `xorm:"DATETIME(8)"`
	Shifttimetag     string    `xorm:"NVARCHAR(510)"`
}

type AnnouncementTag struct {
	Id   int    `xorm:"not null pk INT(4)"`
	Name string `xorm:"NVARCHAR(100)"`
}

type Biddinggroup struct {
	Id   int    `xorm:"not null pk INT(4)"`
	Name string `xorm:"NVARCHAR(100)"`
}

type Swaplaststatuslog struct {
	Swaplaststatuslogid       int64     `xorm:"not null pk BIGINT(8)"`
	Applicationid             []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Substitutiongroupid       []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Swaptype                  string    `xorm:"not null NVARCHAR(32)"`
	Status                    string    `xorm:"not null NVARCHAR(32)"`
	Providedate               time.Time `xorm:"not null DATETIME(8)"`
	Demanddate                time.Time `xorm:"not null DATETIME(8)"`
	Providerid                []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Demanderid                []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Approverid                []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Maturitydate              time.Time `xorm:"not null DATETIME(8)"`
	Lastupdateat              time.Time `xorm:"not null DATETIME(8)"`
	Providername              string    `xorm:"not null NVARCHAR(128)"`
	Demandername              string    `xorm:"not null NVARCHAR(128)"`
	Provideractivityname      string    `xorm:"NVARCHAR(256)"`
	Provideractivitystarttime time.Time `xorm:"DATETIME(8)"`
	Provideractivityendtime   time.Time `xorm:"DATETIME(8)"`
	Demanderactivityname      string    `xorm:"NVARCHAR(256)"`
	Demanderactivitystarttime time.Time `xorm:"DATETIME(8)"`
	Demanderactivityendtime   time.Time `xorm:"DATETIME(8)"`
	Approvername              string    `xorm:"NVARCHAR(128)"`
}

type Backupterm struct {
	Id                 int64     `xorm:"pk BIGINT(8)"`
	Starttime          time.Time `xorm:"not null DATETIME(8)"`
	Finishtime         time.Time `xorm:"not null DATETIME(8)"`
	Parentbackuptermid int64     `xorm:"BIGINT(8)"`
	Background         string    `xorm:"NVARCHAR(510)"`
	Text               string    `xorm:"NVARCHAR(510)"`
	Employeeid         []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Level              int       `xorm:"INT(4)"`
	Workingtotals      int64     `xorm:"BIGINT(8)"`
	Hrdate             time.Time `xorm:"DATETIME(8)"`
}

type Seatconsolidationassignmenttypes struct {
	Seatconsolidationruleid int64  `xorm:"not null pk BIGINT(8)"`
	Assignmenttypeid        []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Fairauditlog struct {
	Id             int64     `xorm:"pk BIGINT(8)"`
	Childclasstype string    `xorm:"not null NVARCHAR(510)"`
	Createat       time.Time `xorm:"DATETIME(8)"`
	Action         string    `xorm:"NVARCHAR(510)"`
	Loginemployee  []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Campaign       []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Csstarttime    time.Time `xorm:"DATETIME(8)"`
	Csfinishtime   time.Time `xorm:"DATETIME(8)"`
	Term           string    `xorm:"NVARCHAR(510)"`
	Termnumber     float32   `xorm:"FLOAT(8)"`
	Employeeid     []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Laborrule      []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Smsqueue struct {
	Id        int       `xorm:"not null pk INT(4)"`
	Phoneno   string    `xorm:"not null VARCHAR(50)"`
	Text      string    `xorm:"not null NVARCHAR(140)"`
	Sendok    int       `xorm:"not null BIT(1)"`
	Sendat    time.Time `xorm:"DATETIME(8)"`
	Createdat time.Time `xorm:"not null DATETIME(8)"`
	Startsend time.Time `xorm:"not null DATETIME(8)"`
	Trycounts int       `xorm:"not null INT(4)"`
	Processat time.Time `xorm:"DATETIME(8)"`
}

type Swapdenypolicy struct {
	Swapdenypolicyid   []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Swapdenypolicyname string `xorm:"not null NVARCHAR(128)"`
	Isavailableforswap int    `xorm:"not null BIT(1)"`
	Sortid             int    `xorm:"INT(4)"`
}

type AnnouncementExtSwapapplication struct {
	Announceid []byte `xorm:"not null pk default newid UNIQUEIDENTIFIER(16)"`
	Logid      []byte `xorm:"not null UNIQUEIDENTIFIER(16)"`
}

type Dailyeventbook struct {
	Id         int       `xorm:"not null pk INT(4)"`
	Targetdate time.Time `xorm:"not null DATETIME(8)"`
	Reader     []byte    `xorm:"VARBINARY(8000)"`
	Note       string    `xorm:"NVARCHAR(8000)"`
}

type Modifiedhistoricaltraffic struct {
	Modifiedhistoricaltrafficid int       `xorm:"not null pk INT(4)"`
	Forecasthistoryid           int       `xorm:"not null INT(4)"`
	Trafficdate                 time.Time `xorm:"not null DATETIME(8)"`
	Aht1                        int       `xorm:"INT(4)"`
	Aht2                        int       `xorm:"INT(4)"`
	Aht3                        int       `xorm:"INT(4)"`
	Aht4                        int       `xorm:"INT(4)"`
	Aht5                        int       `xorm:"INT(4)"`
	Aht6                        int       `xorm:"INT(4)"`
	Aht7                        int       `xorm:"INT(4)"`
	Aht8                        int       `xorm:"INT(4)"`
	Aht9                        int       `xorm:"INT(4)"`
	Aht10                       int       `xorm:"INT(4)"`
	Aht11                       int       `xorm:"INT(4)"`
	Aht12                       int       `xorm:"INT(4)"`
	Aht13                       int       `xorm:"INT(4)"`
	Aht14                       int       `xorm:"INT(4)"`
	Aht15                       int       `xorm:"INT(4)"`
	Aht16                       int       `xorm:"INT(4)"`
	Aht17                       int       `xorm:"INT(4)"`
	Aht18                       int       `xorm:"INT(4)"`
	Aht19                       int       `xorm:"INT(4)"`
	Aht20                       int       `xorm:"INT(4)"`
	Aht21                       int       `xorm:"INT(4)"`
	Aht22                       int       `xorm:"INT(4)"`
	Aht23                       int       `xorm:"INT(4)"`
	Aht24                       int       `xorm:"INT(4)"`
	Aht25                       int       `xorm:"INT(4)"`
	Aht26                       int       `xorm:"INT(4)"`
	Aht27                       int       `xorm:"INT(4)"`
	Aht28                       int       `xorm:"INT(4)"`
	Aht29                       int       `xorm:"INT(4)"`
	Aht30                       int       `xorm:"INT(4)"`
	Aht31                       int       `xorm:"INT(4)"`
	Aht32                       int       `xorm:"INT(4)"`
	Aht33                       int       `xorm:"INT(4)"`
	Aht34                       int       `xorm:"INT(4)"`
	Aht35                       int       `xorm:"INT(4)"`
	Aht36                       int       `xorm:"INT(4)"`
	Aht37                       int       `xorm:"INT(4)"`
	Aht38                       int       `xorm:"INT(4)"`
	Aht39                       int       `xorm:"INT(4)"`
	Aht40                       int       `xorm:"INT(4)"`
	Aht41                       int       `xorm:"INT(4)"`
	Aht42                       int       `xorm:"INT(4)"`
	Aht43                       int       `xorm:"INT(4)"`
	Aht44                       int       `xorm:"INT(4)"`
	Aht45                       int       `xorm:"INT(4)"`
	Aht46                       int       `xorm:"INT(4)"`
	Aht47                       int       `xorm:"INT(4)"`
	Aht48                       int       `xorm:"INT(4)"`
	Aht49                       int       `xorm:"INT(4)"`
	Aht50                       int       `xorm:"INT(4)"`
	Aht51                       int       `xorm:"INT(4)"`
	Aht52                       int       `xorm:"INT(4)"`
	Aht53                       int       `xorm:"INT(4)"`
	Aht54                       int       `xorm:"INT(4)"`
	Aht55                       int       `xorm:"INT(4)"`
	Aht56                       int       `xorm:"INT(4)"`
	Aht57                       int       `xorm:"INT(4)"`
	Aht58                       int       `xorm:"INT(4)"`
	Aht59                       int       `xorm:"INT(4)"`
	Aht60                       int       `xorm:"INT(4)"`
	Aht61                       int       `xorm:"INT(4)"`
	Aht62                       int       `xorm:"INT(4)"`
	Aht63                       int       `xorm:"INT(4)"`
	Aht64                       int       `xorm:"INT(4)"`
	Aht65                       int       `xorm:"INT(4)"`
	Aht66                       int       `xorm:"INT(4)"`
	Aht67                       int       `xorm:"INT(4)"`
	Aht68                       int       `xorm:"INT(4)"`
	Aht69                       int       `xorm:"INT(4)"`
	Aht70                       int       `xorm:"INT(4)"`
	Aht71                       int       `xorm:"INT(4)"`
	Aht72                       int       `xorm:"INT(4)"`
	Aht73                       int       `xorm:"INT(4)"`
	Aht74                       int       `xorm:"INT(4)"`
	Aht75                       int       `xorm:"INT(4)"`
	Aht76                       int       `xorm:"INT(4)"`
	Aht77                       int       `xorm:"INT(4)"`
	Aht78                       int       `xorm:"INT(4)"`
	Aht79                       int       `xorm:"INT(4)"`
	Aht80                       int       `xorm:"INT(4)"`
	Aht81                       int       `xorm:"INT(4)"`
	Aht82                       int       `xorm:"INT(4)"`
	Aht83                       int       `xorm:"INT(4)"`
	Aht84                       int       `xorm:"INT(4)"`
	Aht85                       int       `xorm:"INT(4)"`
	Aht86                       int       `xorm:"INT(4)"`
	Aht87                       int       `xorm:"INT(4)"`
	Aht88                       int       `xorm:"INT(4)"`
	Aht89                       int       `xorm:"INT(4)"`
	Aht90                       int       `xorm:"INT(4)"`
	Aht91                       int       `xorm:"INT(4)"`
	Aht92                       int       `xorm:"INT(4)"`
	Aht93                       int       `xorm:"INT(4)"`
	Aht94                       int       `xorm:"INT(4)"`
	Aht95                       int       `xorm:"INT(4)"`
	Aht96                       int       `xorm:"INT(4)"`
	Cv1                         int       `xorm:"INT(4)"`
	Cv2                         int       `xorm:"INT(4)"`
	Cv3                         int       `xorm:"INT(4)"`
	Cv4                         int       `xorm:"INT(4)"`
	Cv5                         int       `xorm:"INT(4)"`
	Cv6                         int       `xorm:"INT(4)"`
	Cv7                         int       `xorm:"INT(4)"`
	Cv8                         int       `xorm:"INT(4)"`
	Cv9                         int       `xorm:"INT(4)"`
	Cv10                        int       `xorm:"INT(4)"`
	Cv11                        int       `xorm:"INT(4)"`
	Cv12                        int       `xorm:"INT(4)"`
	Cv13                        int       `xorm:"INT(4)"`
	Cv14                        int       `xorm:"INT(4)"`
	Cv15                        int       `xorm:"INT(4)"`
	Cv16                        int       `xorm:"INT(4)"`
	Cv17                        int       `xorm:"INT(4)"`
	Cv18                        int       `xorm:"INT(4)"`
	Cv19                        int       `xorm:"INT(4)"`
	Cv20                        int       `xorm:"INT(4)"`
	Cv21                        int       `xorm:"INT(4)"`
	Cv22                        int       `xorm:"INT(4)"`
	Cv23                        int       `xorm:"INT(4)"`
	Cv24                        int       `xorm:"INT(4)"`
	Cv25                        int       `xorm:"INT(4)"`
	Cv26                        int       `xorm:"INT(4)"`
	Cv27                        int       `xorm:"INT(4)"`
	Cv28                        int       `xorm:"INT(4)"`
	Cv29                        int       `xorm:"INT(4)"`
	Cv30                        int       `xorm:"INT(4)"`
	Cv31                        int       `xorm:"INT(4)"`
	Cv32                        int       `xorm:"INT(4)"`
	Cv33                        int       `xorm:"INT(4)"`
	Cv34                        int       `xorm:"INT(4)"`
	Cv35                        int       `xorm:"INT(4)"`
	Cv36                        int       `xorm:"INT(4)"`
	Cv37                        int       `xorm:"INT(4)"`
	Cv38                        int       `xorm:"INT(4)"`
	Cv39                        int       `xorm:"INT(4)"`
	Cv40                        int       `xorm:"INT(4)"`
	Cv41                        int       `xorm:"INT(4)"`
	Cv42                        int       `xorm:"INT(4)"`
	Cv43                        int       `xorm:"INT(4)"`
	Cv44                        int       `xorm:"INT(4)"`
	Cv45                        int       `xorm:"INT(4)"`
	Cv46                        int       `xorm:"INT(4)"`
	Cv47                        int       `xorm:"INT(4)"`
	Cv48                        int       `xorm:"INT(4)"`
	Cv49                        int       `xorm:"INT(4)"`
	Cv50                        int       `xorm:"INT(4)"`
	Cv51                        int       `xorm:"INT(4)"`
	Cv52                        int       `xorm:"INT(4)"`
	Cv53                        int       `xorm:"INT(4)"`
	Cv54                        int       `xorm:"INT(4)"`
	Cv55                        int       `xorm:"INT(4)"`
	Cv56                        int       `xorm:"INT(4)"`
	Cv57                        int       `xorm:"INT(4)"`
	Cv58                        int       `xorm:"INT(4)"`
	Cv59                        int       `xorm:"INT(4)"`
	Cv60                        int       `xorm:"INT(4)"`
	Cv61                        int       `xorm:"INT(4)"`
	Cv62                        int       `xorm:"INT(4)"`
	Cv63                        int       `xorm:"INT(4)"`
	Cv64                        int       `xorm:"INT(4)"`
	Cv65                        int       `xorm:"INT(4)"`
	Cv66                        int       `xorm:"INT(4)"`
	Cv67                        int       `xorm:"INT(4)"`
	Cv68                        int       `xorm:"INT(4)"`
	Cv69                        int       `xorm:"INT(4)"`
	Cv70                        int       `xorm:"INT(4)"`
	Cv71                        int       `xorm:"INT(4)"`
	Cv72                        int       `xorm:"INT(4)"`
	Cv73                        int       `xorm:"INT(4)"`
	Cv74                        int       `xorm:"INT(4)"`
	Cv75                        int       `xorm:"INT(4)"`
	Cv76                        int       `xorm:"INT(4)"`
	Cv77                        int       `xorm:"INT(4)"`
	Cv78                        int       `xorm:"INT(4)"`
	Cv79                        int       `xorm:"INT(4)"`
	Cv80                        int       `xorm:"INT(4)"`
	Cv81                        int       `xorm:"INT(4)"`
	Cv82                        int       `xorm:"INT(4)"`
	Cv83                        int       `xorm:"INT(4)"`
	Cv84                        int       `xorm:"INT(4)"`
	Cv85                        int       `xorm:"INT(4)"`
	Cv86                        int       `xorm:"INT(4)"`
	Cv87                        int       `xorm:"INT(4)"`
	Cv88                        int       `xorm:"INT(4)"`
	Cv89                        int       `xorm:"INT(4)"`
	Cv90                        int       `xorm:"INT(4)"`
	Cv91                        int       `xorm:"INT(4)"`
	Cv92                        int       `xorm:"INT(4)"`
	Cv93                        int       `xorm:"INT(4)"`
	Cv94                        int       `xorm:"INT(4)"`
	Cv95                        int       `xorm:"INT(4)"`
	Cv96                        int       `xorm:"INT(4)"`
}

type Groupdistribution struct {
	Groupdistributionid int64  `xorm:"not null pk BIGINT(8)"`
	Laborruleid         []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Assignmenttypeid    []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Weight              string `xorm:"NVARCHAR(510)"`
	Quotaenforcement    int    `xorm:"INT(4)"`
	Tag                 string `xorm:"NVARCHAR(510)"`
}

type Timeoffgroup struct {
	Timeoffgroupid        []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name                  string `xorm:"not null NVARCHAR(510)"`
	Workyearlimit         int    `xorm:"INT(4)"`
	Applyovertime         int    `xorm:"INT(4)"`
	Monthlimit            int    `xorm:"INT(4)"`
	Maximumcontinuousdays int    `xorm:"INT(4)"`
}

type AnnouncementExtShiftmodified struct {
	Announceid []byte `xorm:"not null pk default newid UNIQUEIDENTIFIER(16)"`
}

type Dailyeventbookreader struct {
	Id               int       `xorm:"not null pk INT(4)"`
	Dailyeventbookid int       `xorm:"not null INT(4)"`
	Employeeid       []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Readdate         time.Time `xorm:"not null default 'getdate' DATETIME(8)"`
}

type Customgroupemployees struct {
	Customemployeegroupid int    `xorm:"not null pk INT(4)"`
	Elt                   []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Groupdistributiondailycounter struct {
	Groupdistributiondailycounterid int64     `xorm:"not null pk BIGINT(8)"`
	Laborruleid                     []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Assignmenttypeid                []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Quota                           float32   `xorm:"FLOAT(8)"`
	Date                            time.Time `xorm:"not null DATETIME(8)"`
}

type Campaignorg struct {
	Campaignid     []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Organizationid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Swapapplicationapproval struct {
	Swapapplicationapprovalid []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Swapapplicationid         []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Managerid                 []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Managerinfo               string    `xorm:"not null NVARCHAR(128)"`
	Employeeid                []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Approveddate              time.Time `xorm:"not null DATETIME(8)"`
	Status                    string    `xorm:"not null NVARCHAR(32)"`
	Comment                   string    `xorm:"NVARCHAR(510)"`
}

type AnnouncementExtAbsence struct {
	Announceid []byte `xorm:"not null pk default newid UNIQUEIDENTIFIER(16)"`
}

type Modifiedforecastrecord struct {
	Modifiedforecastrecordid int       `xorm:"not null pk INT(4)"`
	Forecastid               int       `xorm:"not null INT(4)"`
	Createtime               time.Time `xorm:"not null DATETIME(8)"`
	Datetype                 int       `xorm:"not null INT(4)"`
	Adjustitem               int       `xorm:"not null BIT(1)"`
	Targetdate               string    `xorm:"not null NVARCHAR(2000)"`
	Adjustmenttype           int       `xorm:"not null INT(4)"`
	Value                    int       `xorm:"INT(4)"`
	Remark                   string    `xorm:"NVARCHAR(600)"`
	Creator                  []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Batchindex               int       `xorm:"not null INT(4)"`
}

type Event struct {
	Id          int       `xorm:"not null pk INT(4)"`
	Eventbookid int       `xorm:"not null INT(4)"`
	Startat     time.Time `xorm:"not null DATETIME(8)"`
	Endat       time.Time `xorm:"not null DATETIME(8)"`
	Stopat      time.Time `xorm:"not null DATETIME(8)"`
	Eventlevel  int       `xorm:"not null INT(4)"`
	Eventtypeid int       `xorm:"not null INT(4)"`
	Happenwith  string    `xorm:"VARCHAR(1024)"`
	Eventreason string    `xorm:"NVARCHAR(400)"`
	Eventhandle string    `xorm:"NVARCHAR(400)"`
	Note        string    `xorm:"NVARCHAR(400)"`
	Judgment    string    `xorm:"NVARCHAR(400)"`
	Createdby   []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Createdat   time.Time `xorm:"not null default 'getdate' DATETIME(8)"`
	Iscompleted int       `xorm:"default 0 BIT(1)"`
	Isscored    int       `xorm:"not null default 0 BIT(1)"`
	Score       int       `xorm:"not null default 0 INT(4)"`
	Isforfeited int       `xorm:"not null default 0 BIT(1)"`
	Forfeit     int       `xorm:"not null default 0 INT(4)"`
	Recordid    string    `xorm:"VARCHAR(50)"`
	Recordat    time.Time `xorm:"DATETIME(8)"`
	Projectid   string    `xorm:"VARCHAR(50)"`
	Islogin     int       `xorm:"BIT(1)"`
}

type Timeoffgrouppolicy struct {
	Timeoffgrouppolicyid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name                 string `xorm:"not null NVARCHAR(510)"`
	Isavailable          int    `xorm:"BIT(1)"`
}

type Agentwebauditlog struct {
	Id         int64     `xorm:"pk BIGINT(8)"`
	Time       time.Time `xorm:"not null default 'getdate' DATETIME(8)"`
	Issuccess  int       `xorm:"not null BIT(1)"`
	Loginname  string    `xorm:"not null NVARCHAR(100)"`
	Employeeid string    `xorm:"not null NVARCHAR(100)"`
	Ip         string    `xorm:"not null NVARCHAR(50)"`
	Info       string    `xorm:"not null NVARCHAR(200)"`
	Attemppass string    `xorm:"not null NVARCHAR(200)"`
}

type Termgaprule struct {
	Termgapruleid   int64  `xorm:"not null pk BIGINT(8)"`
	Assignmentnamea string `xorm:"NVARCHAR(510)"`
	Assignmentnameb string `xorm:"NVARCHAR(510)"`
	Laborhour       int64  `xorm:"BIGINT(8)"`
	Laborruleid     []byte `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Modifiedhistoricalrecord struct {
	Modifiedhistoricalrecordid int       `xorm:"not null pk INT(4)"`
	Forecasthistoryid          int       `xorm:"not null INT(4)"`
	Createtime                 time.Time `xorm:"not null DATETIME(8)"`
	Datetype                   int       `xorm:"not null INT(4)"`
	Adjustitem                 int       `xorm:"not null BIT(1)"`
	Targetdate                 string    `xorm:"not null NVARCHAR(2000)"`
	Adjustmenttype             int       `xorm:"not null INT(4)"`
	Value                      int       `xorm:"INT(4)"`
	Remark                     string    `xorm:"NVARCHAR(600)"`
	Creator                    []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Batchindex                 int       `xorm:"not null INT(4)"`
}

type Acdqueue struct {
	Acdqueueid     []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name           string `xorm:"NVARCHAR(510)"`
	Description    string `xorm:"NVARCHAR(510)"`
	Splitid        string `xorm:"NVARCHAR(510)"`
	Servicequeueid []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Acdid          []byte `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Swapapplication struct {
	Swapapplicationid   []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Applicationdate     time.Time `xorm:"not null DATETIME(8)"`
	Employeeaid         []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Employeeainfo       string    `xorm:"not null NVARCHAR(128)"`
	Datea               time.Time `xorm:"not null DATETIME(8)"`
	Assignmentaid       int64     `xorm:"BIGINT(8)"`
	Assignmentainfo     string    `xorm:"NVARCHAR(510)"`
	Descriptiona        string    `xorm:"NVARCHAR(510)"`
	Employeebid         []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Employeebinfo       string    `xorm:"not null NVARCHAR(128)"`
	Dateb               time.Time `xorm:"not null DATETIME(8)"`
	Assignmentbid       int64     `xorm:"BIGINT(8)"`
	Assignmentbinfo     string    `xorm:"NVARCHAR(510)"`
	Descriptionb        string    `xorm:"NVARCHAR(510)"`
	Status              string    `xorm:"not null NVARCHAR(32)"`
	Completedate        time.Time `xorm:"DATETIME(8)"`
	Maturitydate        time.Time `xorm:"not null DATETIME(8)"`
	Swaptype            string    `xorm:"not null NVARCHAR(32)"`
	Substitutiongroupid []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Warningmessagea     string    `xorm:"NVARCHAR(-1)"`
	Warningmessageb     string    `xorm:"NVARCHAR(-1)"`
	Autoagreemessagea   string    `xorm:"NVARCHAR(-1)"`
	Autoagreemessageb   string    `xorm:"NVARCHAR(-1)"`
	Iscutshifta         int       `xorm:"not null default 0 BIT(1)"`
	Iscutshiftb         int       `xorm:"not null default 0 BIT(1)"`
	Version             time.Time `xorm:"TIMESTAMP(8)"`
}

type Moduleresource struct {
	Id     []byte `xorm:"not null pk unique UNIQUEIDENTIFIER(16)"`
	Name   string `xorm:"not null NVARCHAR(510)"`
	Parent []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Enname string `xorm:"NVARCHAR(510)"`
}

type Termtimesrule struct {
	Termtimesruleid int64  `xorm:"not null pk BIGINT(8)"`
	Shifttimetag    string `xorm:"NVARCHAR(510)"`
	Maxtimes        int    `xorm:"INT(4)"`
	Laborruleid     []byte `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Cssq struct {
	Campaignscheduleid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Minonlines         int    `xorm:"INT(4)"`
	Servicequeueid     []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Swapagreepolicy struct {
	Swapagreepolicyid   []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Swapagreepolicyname string `xorm:"not null NVARCHAR(128)"`
	Isavailableforswap  int    `xorm:"not null BIT(1)"`
	Sortid              int    `xorm:"INT(4)"`
}

type Authrole struct {
	Roleid   []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Rolename string `xorm:"NVARCHAR(510)"`
}

type Adherenceevent struct {
	Adherenceeventid int64     `xorm:"not null pk BIGINT(8)"`
	Timestamp        time.Time `xorm:"not null DATETIME(8)"`
	Event            string    `xorm:"NVARCHAR(510)"`
	Starttime        time.Time `xorm:"not null index(IX_1) DATETIME(8)"`
	Finishtime       time.Time `xorm:"not null index(IX_1) DATETIME(8)"`
	Employeeid       []byte    `xorm:"not null index(IX_1) UNIQUEIDENTIFIER(16)"`
	Assigner         string    `xorm:"NVARCHAR(510)"`
	Reason           string    `xorm:"NVARCHAR(510)"`
}

type TAgentStateSwitch struct {
	Id         int       `xorm:"not null pk INT(4)"`
	Agentid    string    `xorm:"VARCHAR(50)"`
	Extno      string    `xorm:"VARCHAR(50)"`
	State      string    `xorm:"VARCHAR(50)"`
	Workmode   int       `xorm:"INT(4)"`
	Eventtime  time.Time `xorm:"DATETIME(8)"`
	Inserttime time.Time `xorm:"DATETIME(8)"`
}

type Csorg struct {
	Campaignscheduleid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Organizationid     []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Swapagenttimeoff struct {
	Swappostid   []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Appilerid    []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Reappilerid  []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Dateappier   time.Time `xorm:"not null DATETIME(8)"`
	Datereappier time.Time `xorm:"not null DATETIME(8)"`
	Poststart    time.Time `xorm:"not null DATETIME(8)"`
	Postend      time.Time `xorm:"DATETIME(8)"`
	Swapstate    int       `xorm:"not null INT(4)"`
}

type Mailqueue struct {
	Id           int       `xorm:"not null pk INT(4)"`
	Emailaddress string    `xorm:"not null VARCHAR(5000)"`
	Subject      string    `xorm:"not null NVARCHAR(100)"`
	Text         string    `xorm:"not null NVARCHAR(-1)"`
	Sendok       int       `xorm:"not null BIT(1)"`
	Sendat       time.Time `xorm:"DATETIME(8)"`
	Createdat    time.Time `xorm:"not null DATETIME(8)"`
	Startsend    time.Time `xorm:"not null DATETIME(8)"`
	Trycounts    int       `xorm:"not null INT(4)"`
	Processat    time.Time `xorm:"DATETIME(8)"`
	Contenttype  string    `xorm:"NVARCHAR(40)"`
	Postername   string    `xorm:"NVARCHAR(40)"`
}

type Authfunction struct {
	Functionkey string `xorm:"not null pk NVARCHAR(510)"`
	Displayname string `xorm:"not null unique NVARCHAR(510)"`
	Category    string `xorm:"not null NVARCHAR(510)"`
}

type Agentstatustype struct {
	Agentstatustypecode string `xorm:"not null pk NVARCHAR(510)"`
	Statusname          string `xorm:"not null NVARCHAR(64)"`
	Onservice           int    `xorm:"not null BIT(1)"`
	Alerttimeoutsecond  int    `xorm:"not null INT(4)"`
	Alerttimeoutsecond2 int    `xorm:"not null INT(4)"`
	Imagebyte           []byte `xorm:"VARBINARY(-1)"`
	Islogout            int    `xorm:"not null BIT(1)"`
	Islogin             int    `xorm:"not null BIT(1)"`
}

type Substitutiongroupswapdenypolicy struct {
	Substitutiongroupid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Swapdenypolicyid    []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Settleshiftsdeviation struct {
	Id             int64     `xorm:"pk BIGINT(8)"`
	Childclasstype string    `xorm:"not null NVARCHAR(64)"`
	Starttime      time.Time `xorm:"not null DATETIME(8)"`
	Finishtime     time.Time `xorm:"not null DATETIME(8)"`
	Originalterm   string    `xorm:"NVARCHAR(510)"`
	Combineto      string    `xorm:"NVARCHAR(510)"`
	Count          float32   `xorm:"FLOAT(8)"`
	Deviation      float32   `xorm:"FLOAT(8)"`
	Settled        int       `xorm:"BIT(1)"`
	Employeeid     []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Laborruleid    []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Campaignid     []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Summons        float32   `xorm:"FLOAT(8)"`
	Remark         string    `xorm:"NVARCHAR(510)"`
	Createat       time.Time `xorm:"DATETIME(8)"`
	Loginemployee  []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Surplus        int       `xorm:"BIT(1)"`
}

type AnnouncementOrg struct {
	Announceid        []byte    `xorm:"not null pk default newid UNIQUEIDENTIFIER(16)"`
	Posttime          time.Time `xorm:"not null default 'getdate' DATETIME(8)"`
	Poster            string    `xorm:"not null NVARCHAR(1020)"`
	Posterid          []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Receiverorgid     []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Title             string    `xorm:"not null default 'N'No Title'' NVARCHAR(100)"`
	Contenttext       string    `xorm:"not null NVARCHAR(-1)"`
	Type              int       `xorm:"not null INT(4)"`
	Tag               string    `xorm:"NVARCHAR(100)"`
	Priority          int       `xorm:"BIT(1)"`
	Announcetimestart time.Time `xorm:"not null DATETIME(8)"`
	Announcetimeend   time.Time `xorm:"not null DATETIME(8)"`
	Contenttype       string    `xorm:"not null default ''Web'' NVARCHAR(100)"`
	Isdeleted         int       `xorm:"BIT(1)"`
}

type Substitutiongroupswapagreepolicy struct {
	Substitutiongroupid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Swapagreepolicyid   []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type AuditlogExt struct {
	Id      int64     `xorm:"pk BIGINT(8)"`
	Logtime time.Time `xorm:"not null DATETIME(8)"`
	Agentid string    `xorm:"NVARCHAR(510)"`
	Action  string    `xorm:"NVARCHAR(510)"`
}

type Auditemployeeskill struct {
	Fid         int       `xorm:"not null INT(4)"`
	Ftime       time.Time `xorm:"DATETIME(8)"`
	Femployeeid []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Fname       string    `xorm:"NVARCHAR(100)"`
	Faction     string    `xorm:"NVARCHAR(510)"`
	Frecord     string    `xorm:"NVARCHAR(510)"`
}

type Substitutiongroupseswappolicy struct {
	Substitutiongroupid  []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Subeventswappolicyid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Authroleemployee struct {
	Roleid     []byte `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Employeeid []byte `xorm:"not null UNIQUEIDENTIFIER(16)"`
}

type Auditagentacdidcolor struct {
	Fid     int       `xorm:"not null pk INT(4)"`
	Ftime   time.Time `xorm:"DATETIME(8)"`
	Fname   string    `xorm:"NVARCHAR(100)"`
	Faction string    `xorm:"NVARCHAR(100)"`
	Frecord string    `xorm:"NVARCHAR(100)"`
}

type Assignmenttypedailycounter struct {
	Assignmenttypedailycounterid int64     `xorm:"not null pk BIGINT(8)"`
	Assignmenttypeid             []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Max                          float32   `xorm:"FLOAT(8)"`
	Date                         time.Time `xorm:"not null DATETIME(8)"`
}

type AnnouncementPersonal struct {
	Announceid         []byte    `xorm:"not null pk default newid UNIQUEIDENTIFIER(16)"`
	Posttime           time.Time `xorm:"not null default 'getdate' DATETIME(8)"`
	Poster             string    `xorm:"not null NVARCHAR(1020)"`
	Posterid           []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Receiveremployeeid []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Title              string    `xorm:"not null default 'N'No Title'' NVARCHAR(100)"`
	Contenttext        string    `xorm:"not null NVARCHAR(-1)"`
	Type               int       `xorm:"not null INT(4)"`
	Tag                string    `xorm:"NVARCHAR(100)"`
	Priority           int       `xorm:"BIT(1)"`
	Isread             int       `xorm:"not null default 0 BIT(1)"`
	Isshow             int       `xorm:"not null default 1 BIT(1)"`
	Contenttype        string    `xorm:"default ''Web'' NVARCHAR(100)"`
	Isdeleted          int       `xorm:"BIT(1)"`
}

type Config struct {
	Keyname  string `xorm:"not null pk NVARCHAR(400)"`
	Keyvalue string `xorm:"not null NVARCHAR(400)"`
}

type Agentsortsetting struct {
	Biddingannouncementid []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Employeeid            []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Sortdatedetail        string    `xorm:"not null NVARCHAR(-1)"`
	Timeoffdate           string    `xorm:"NVARCHAR(510)"`
	Lastmodifyat          time.Time `xorm:"DATETIME(8)"`
}

type Substitutiongroupapprovalmanager struct {
	Substitutiongroupid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Managerid           []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Managertype         string `xorm:"not null pk NVARCHAR(32)"`
}

type Agentscoretype struct {
	Id        int     `xorm:"not null pk INT(4)"`
	Name      string  `xorm:"not null NVARCHAR(100)"`
	Scorerule string  `xorm:"not null NVARCHAR(100)"`
	Minvalue  float32 `xorm:"not null default 0 FLOAT(8)"`
	Maxvalue  float32 `xorm:"not null default 0 FLOAT(8)"`
	Inused    int     `xorm:"default 1 BIT(1)"`
}

type Agentpreference struct {
	Employeeid          []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Legalholiday        int    `xorm:"INT(4)"`
	Weekdayperfect      string `xorm:"NVARCHAR(510)"`
	Monthdayperfect     string `xorm:"NVARCHAR(510)"`
	Timeintervelperfect string `xorm:"NVARCHAR(1200)"`
}

type Offdutegroup struct {
	Offdutegroupid   []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name             string `xorm:"VARCHAR(50)"`
	Replacetermname  string `xorm:"VARCHAR(50)"`
	Amtermname       string `xorm:"VARCHAR(50)"`
	Pmtermname       string `xorm:"VARCHAR(50)"`
	Expriebeforedays int    `xorm:"INT(4)"`
	Offdutetimesum   int    `xorm:"INT(4)"`
}

type Agentacdid struct {
	Loginid    int    `xorm:"not null pk INT(4)"`
	Employeeid []byte `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Auditlog struct {
	Id      int64     `xorm:"pk BIGINT(8)"`
	Logtime time.Time `xorm:"not null DATETIME(8)"`
	Agentid string    `xorm:"NVARCHAR(510)"`
	Action  string    `xorm:"NVARCHAR(510)"`
	Ip      string    `xorm:"NVARCHAR(100)"`
}

type Csagent struct {
	Csagentid                  int       `xorm:"not null pk INT(4)"`
	Employeeid                 []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Campaign                   []byte    `xorm:"index(IX_Attendance) UNIQUEIDENTIFIER(16)"`
	Start                      time.Time `xorm:"index(IX_Attendance) DATETIME(8)"`
	Finsh                      time.Time `xorm:"index(IX_Attendance) DATETIME(8)"`
	Maxovertimecurrent         int64     `xorm:"BIGINT(8)"`
	Maxovertimethreshold       int64     `xorm:"BIGINT(8)"`
	Maxshrinkedcurrent         int64     `xorm:"BIGINT(8)"`
	Maxshrinkedthreshold       int64     `xorm:"BIGINT(8)"`
	Mcdo                       int       `xorm:"INT(4)"`
	Mcwd                       int       `xorm:"INT(4)"`
	Minidlegap                 int64     `xorm:"BIGINT(8)"`
	Stddailylaborhour          int64     `xorm:"BIGINT(8)"`
	Maxlaborhour               int64     `xorm:"BIGINT(8)"`
	Minlaborhour               int64     `xorm:"BIGINT(8)"`
	Laborhourtotals            int64     `xorm:"BIGINT(8)"`
	Maxswaptimes               int       `xorm:"INT(4)"`
	Schedulingpayload          int       `xorm:"unique INT(4)"`
	Amountdayoff               int       `xorm:"INT(4)"`
	Score                      float32   `xorm:"FLOAT(8)"`
	Systemaccumulate           int       `xorm:"BIT(1)"`
	Add1dayoffeachsaturdayincs int       `xorm:"BIT(1)"`
	Add1dayoffeachsundayincs   int       `xorm:"BIT(1)"`
	Add1dayoffeachholidayincs  int       `xorm:"BIT(1)"`
	Specifiednumberofdays      int       `xorm:"INT(4)"`
	Holidayshiftrule           int       `xorm:"INT(4)"`
	Maxfwtimes                 int       `xorm:"INT(4)"`
	Minfwtimes                 int       `xorm:"INT(4)"`
	Maxpwtimes                 int       `xorm:"INT(4)"`
	Minpwtimes                 int       `xorm:"INT(4)"`
	Joinedlaborruleid          []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Subeventswappolicy struct {
	Subeventswappolicyid   []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Subeventswappolicyname string `xorm:"not null NVARCHAR(64)"`
	Sortid                 int    `xorm:"INT(4)"`
}

type Employeebiddingbase struct {
	Employeeid   []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Dayoffvalue  int    `xorm:"not null INT(4)"`
	Timeoffvalue int    `xorm:"not null INT(4)"`
}

type Occupation struct {
	Occupationid   int64     `xorm:"not null pk BIGINT(8)"`
	Childclasstype string    `xorm:"not null NVARCHAR(64)"`
	Starttime      time.Time `xorm:"not null DATETIME(8)"`
	Finishtime     time.Time `xorm:"not null DATETIME(8)"`
	Seatboxid      []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Category       string    `xorm:"NVARCHAR(510)"`
	Description    string    `xorm:"NVARCHAR(510)"`
}

type Termlog struct {
	Id              int64     `xorm:"pk BIGINT(8)"`
	Alterat         time.Time `xorm:"not null DATETIME(8)"`
	Sourceid        int64     `xorm:"BIGINT(8)"`
	Alteremployeeid []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Employeeid      []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Action          string    `xorm:"NVARCHAR(510)"`
	Type            string    `xorm:"NVARCHAR(510)"`
	Name            string    `xorm:"NVARCHAR(510)"`
	Category        string    `xorm:"NVARCHAR(510)"`
	Oldtime         string    `xorm:"NVARCHAR(510)"`
	Newtime         string    `xorm:"NVARCHAR(510)"`
	Remark          string    `xorm:"NVARCHAR(510)"`
	Sent            int       `xorm:"BIT(1)"`
}

type Forecasttraffic struct {
	Forecasttrafficid int       `xorm:"not null pk INT(4)"`
	Trafficdate       time.Time `xorm:"not null DATETIME(8)"`
	Forecastid        int       `xorm:"not null INT(4)"`
	Cv1               int       `xorm:"INT(4)"`
	Cv2               int       `xorm:"INT(4)"`
	Cv3               int       `xorm:"INT(4)"`
	Cv4               int       `xorm:"INT(4)"`
	Cv5               int       `xorm:"INT(4)"`
	Cv6               int       `xorm:"INT(4)"`
	Cv7               int       `xorm:"INT(4)"`
	Cv8               int       `xorm:"INT(4)"`
	Cv9               int       `xorm:"INT(4)"`
	Cv10              int       `xorm:"INT(4)"`
	Cv11              int       `xorm:"INT(4)"`
	Cv12              int       `xorm:"INT(4)"`
	Cv13              int       `xorm:"INT(4)"`
	Cv14              int       `xorm:"INT(4)"`
	Cv15              int       `xorm:"INT(4)"`
	Cv16              int       `xorm:"INT(4)"`
	Cv17              int       `xorm:"INT(4)"`
	Cv18              int       `xorm:"INT(4)"`
	Cv19              int       `xorm:"INT(4)"`
	Cv20              int       `xorm:"INT(4)"`
	Cv21              int       `xorm:"INT(4)"`
	Cv22              int       `xorm:"INT(4)"`
	Cv23              int       `xorm:"INT(4)"`
	Cv24              int       `xorm:"INT(4)"`
	Cv25              int       `xorm:"INT(4)"`
	Cv26              int       `xorm:"INT(4)"`
	Cv27              int       `xorm:"INT(4)"`
	Cv28              int       `xorm:"INT(4)"`
	Cv29              int       `xorm:"INT(4)"`
	Cv30              int       `xorm:"INT(4)"`
	Cv31              int       `xorm:"INT(4)"`
	Cv32              int       `xorm:"INT(4)"`
	Cv33              int       `xorm:"INT(4)"`
	Cv34              int       `xorm:"INT(4)"`
	Cv35              int       `xorm:"INT(4)"`
	Cv36              int       `xorm:"INT(4)"`
	Cv37              int       `xorm:"INT(4)"`
	Cv38              int       `xorm:"INT(4)"`
	Cv39              int       `xorm:"INT(4)"`
	Cv40              int       `xorm:"INT(4)"`
	Cv41              int       `xorm:"INT(4)"`
	Cv42              int       `xorm:"INT(4)"`
	Cv43              int       `xorm:"INT(4)"`
	Cv44              int       `xorm:"INT(4)"`
	Cv45              int       `xorm:"INT(4)"`
	Cv46              int       `xorm:"INT(4)"`
	Cv47              int       `xorm:"INT(4)"`
	Cv48              int       `xorm:"INT(4)"`
	Cv49              int       `xorm:"INT(4)"`
	Cv50              int       `xorm:"INT(4)"`
	Cv51              int       `xorm:"INT(4)"`
	Cv52              int       `xorm:"INT(4)"`
	Cv53              int       `xorm:"INT(4)"`
	Cv54              int       `xorm:"INT(4)"`
	Cv55              int       `xorm:"INT(4)"`
	Cv56              int       `xorm:"INT(4)"`
	Cv57              int       `xorm:"INT(4)"`
	Cv58              int       `xorm:"INT(4)"`
	Cv59              int       `xorm:"INT(4)"`
	Cv60              int       `xorm:"INT(4)"`
	Cv61              int       `xorm:"INT(4)"`
	Cv62              int       `xorm:"INT(4)"`
	Cv63              int       `xorm:"INT(4)"`
	Cv64              int       `xorm:"INT(4)"`
	Cv65              int       `xorm:"INT(4)"`
	Cv66              int       `xorm:"INT(4)"`
	Cv67              int       `xorm:"INT(4)"`
	Cv68              int       `xorm:"INT(4)"`
	Cv69              int       `xorm:"INT(4)"`
	Cv70              int       `xorm:"INT(4)"`
	Cv71              int       `xorm:"INT(4)"`
	Cv72              int       `xorm:"INT(4)"`
	Cv73              int       `xorm:"INT(4)"`
	Cv74              int       `xorm:"INT(4)"`
	Cv75              int       `xorm:"INT(4)"`
	Cv76              int       `xorm:"INT(4)"`
	Cv77              int       `xorm:"INT(4)"`
	Cv78              int       `xorm:"INT(4)"`
	Cv79              int       `xorm:"INT(4)"`
	Cv80              int       `xorm:"INT(4)"`
	Cv81              int       `xorm:"INT(4)"`
	Cv82              int       `xorm:"INT(4)"`
	Cv83              int       `xorm:"INT(4)"`
	Cv84              int       `xorm:"INT(4)"`
	Cv85              int       `xorm:"INT(4)"`
	Cv86              int       `xorm:"INT(4)"`
	Cv87              int       `xorm:"INT(4)"`
	Cv88              int       `xorm:"INT(4)"`
	Cv89              int       `xorm:"INT(4)"`
	Cv90              int       `xorm:"INT(4)"`
	Cv91              int       `xorm:"INT(4)"`
	Cv92              int       `xorm:"INT(4)"`
	Cv93              int       `xorm:"INT(4)"`
	Cv94              int       `xorm:"INT(4)"`
	Cv95              int       `xorm:"INT(4)"`
	Cv96              int       `xorm:"INT(4)"`
	Aht1              int       `xorm:"INT(4)"`
	Aht2              int       `xorm:"INT(4)"`
	Aht3              int       `xorm:"INT(4)"`
	Aht4              int       `xorm:"INT(4)"`
	Aht5              int       `xorm:"INT(4)"`
	Aht6              int       `xorm:"INT(4)"`
	Aht7              int       `xorm:"INT(4)"`
	Aht8              int       `xorm:"INT(4)"`
	Aht9              int       `xorm:"INT(4)"`
	Aht10             int       `xorm:"INT(4)"`
	Aht11             int       `xorm:"INT(4)"`
	Aht12             int       `xorm:"INT(4)"`
	Aht13             int       `xorm:"INT(4)"`
	Aht14             int       `xorm:"INT(4)"`
	Aht15             int       `xorm:"INT(4)"`
	Aht16             int       `xorm:"INT(4)"`
	Aht17             int       `xorm:"INT(4)"`
	Aht18             int       `xorm:"INT(4)"`
	Aht19             int       `xorm:"INT(4)"`
	Aht20             int       `xorm:"INT(4)"`
	Aht21             int       `xorm:"INT(4)"`
	Aht22             int       `xorm:"INT(4)"`
	Aht23             int       `xorm:"INT(4)"`
	Aht24             int       `xorm:"INT(4)"`
	Aht25             int       `xorm:"INT(4)"`
	Aht26             int       `xorm:"INT(4)"`
	Aht27             int       `xorm:"INT(4)"`
	Aht28             int       `xorm:"INT(4)"`
	Aht29             int       `xorm:"INT(4)"`
	Aht30             int       `xorm:"INT(4)"`
	Aht31             int       `xorm:"INT(4)"`
	Aht32             int       `xorm:"INT(4)"`
	Aht33             int       `xorm:"INT(4)"`
	Aht34             int       `xorm:"INT(4)"`
	Aht35             int       `xorm:"INT(4)"`
	Aht36             int       `xorm:"INT(4)"`
	Aht37             int       `xorm:"INT(4)"`
	Aht38             int       `xorm:"INT(4)"`
	Aht39             int       `xorm:"INT(4)"`
	Aht40             int       `xorm:"INT(4)"`
	Aht41             int       `xorm:"INT(4)"`
	Aht42             int       `xorm:"INT(4)"`
	Aht43             int       `xorm:"INT(4)"`
	Aht44             int       `xorm:"INT(4)"`
	Aht45             int       `xorm:"INT(4)"`
	Aht46             int       `xorm:"INT(4)"`
	Aht47             int       `xorm:"INT(4)"`
	Aht48             int       `xorm:"INT(4)"`
	Aht49             int       `xorm:"INT(4)"`
	Aht50             int       `xorm:"INT(4)"`
	Aht51             int       `xorm:"INT(4)"`
	Aht52             int       `xorm:"INT(4)"`
	Aht53             int       `xorm:"INT(4)"`
	Aht54             int       `xorm:"INT(4)"`
	Aht55             int       `xorm:"INT(4)"`
	Aht56             int       `xorm:"INT(4)"`
	Aht57             int       `xorm:"INT(4)"`
	Aht58             int       `xorm:"INT(4)"`
	Aht59             int       `xorm:"INT(4)"`
	Aht60             int       `xorm:"INT(4)"`
	Aht61             int       `xorm:"INT(4)"`
	Aht62             int       `xorm:"INT(4)"`
	Aht63             int       `xorm:"INT(4)"`
	Aht64             int       `xorm:"INT(4)"`
	Aht65             int       `xorm:"INT(4)"`
	Aht66             int       `xorm:"INT(4)"`
	Aht67             int       `xorm:"INT(4)"`
	Aht68             int       `xorm:"INT(4)"`
	Aht69             int       `xorm:"INT(4)"`
	Aht70             int       `xorm:"INT(4)"`
	Aht71             int       `xorm:"INT(4)"`
	Aht72             int       `xorm:"INT(4)"`
	Aht73             int       `xorm:"INT(4)"`
	Aht74             int       `xorm:"INT(4)"`
	Aht75             int       `xorm:"INT(4)"`
	Aht76             int       `xorm:"INT(4)"`
	Aht77             int       `xorm:"INT(4)"`
	Aht78             int       `xorm:"INT(4)"`
	Aht79             int       `xorm:"INT(4)"`
	Aht80             int       `xorm:"INT(4)"`
	Aht81             int       `xorm:"INT(4)"`
	Aht82             int       `xorm:"INT(4)"`
	Aht83             int       `xorm:"INT(4)"`
	Aht84             int       `xorm:"INT(4)"`
	Aht85             int       `xorm:"INT(4)"`
	Aht86             int       `xorm:"INT(4)"`
	Aht87             int       `xorm:"INT(4)"`
	Aht88             int       `xorm:"INT(4)"`
	Aht89             int       `xorm:"INT(4)"`
	Aht90             int       `xorm:"INT(4)"`
	Aht91             int       `xorm:"INT(4)"`
	Aht92             int       `xorm:"INT(4)"`
	Aht93             int       `xorm:"INT(4)"`
	Aht94             int       `xorm:"INT(4)"`
	Aht95             int       `xorm:"INT(4)"`
	Aht96             int       `xorm:"INT(4)"`
}

type Employee struct {
	Employeeid      []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name            string    `xorm:"NVARCHAR(510)"`
	Name2           string    `xorm:"NVARCHAR(510)"`
	Enrollmentdate  time.Time `xorm:"DATETIME(8)"`
	Leavingdate     time.Time `xorm:"DATETIME(8)"`
	Employeetag1    int       `xorm:"BIT(1)"`
	Employeetag2    int       `xorm:"BIT(1)"`
	Agentid         string    `xorm:"unique NVARCHAR(510)"`
	Isagent         int       `xorm:"BIT(1)"`
	Rank            int       `xorm:"INT(4)"`
	Employeetype1   string    `xorm:"NVARCHAR(510)"`
	Employeetype2   string    `xorm:"NVARCHAR(510)"`
	Employeetype3   string    `xorm:"NVARCHAR(510)"`
	Language        string    `xorm:"NVARCHAR(510)"`
	Address1        string    `xorm:"NVARCHAR(510)"`
	Address2        string    `xorm:"NVARCHAR(510)"`
	Email           string    `xorm:"NVARCHAR(510)"`
	Mobile          string    `xorm:"NVARCHAR(510)"`
	Productivity    float32   `xorm:"FLOAT(8)"`
	Password        string    `xorm:"NVARCHAR(510)"`
	Laborruleid     []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Organizationid  []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Certificationid string    `xorm:"NVARCHAR(80)"`
	Ondutydate      time.Time `xorm:"DATETIME(8)"`
	Homephone1      string    `xorm:"NVARCHAR(510)"`
	Homephone2      string    `xorm:"NVARCHAR(510)"`
	Iscanworked     int       `xorm:"BIT(1)"`
	Swapgroupid     []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Offdutegroupid  []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Timeoffgroupid  []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Biddinggroupid  int       `xorm:"INT(4)"`
}

type Acdqueuetraffic struct {
	Acdqueuetrafficid int64     `xorm:"not null pk BIGINT(8)"`
	Trafficdate       time.Time `xorm:"DATETIME(8)"`
	Acdqueueid        []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	V0                float32   `xorm:"FLOAT(8)"`
	V1                float32   `xorm:"FLOAT(8)"`
	V2                float32   `xorm:"FLOAT(8)"`
	V3                float32   `xorm:"FLOAT(8)"`
	V4                float32   `xorm:"FLOAT(8)"`
	V5                float32   `xorm:"FLOAT(8)"`
	V6                float32   `xorm:"FLOAT(8)"`
	V7                float32   `xorm:"FLOAT(8)"`
	V8                float32   `xorm:"FLOAT(8)"`
	V9                float32   `xorm:"FLOAT(8)"`
	V10               float32   `xorm:"FLOAT(8)"`
	V11               float32   `xorm:"FLOAT(8)"`
	V12               float32   `xorm:"FLOAT(8)"`
	V13               float32   `xorm:"FLOAT(8)"`
	V14               float32   `xorm:"FLOAT(8)"`
	V15               float32   `xorm:"FLOAT(8)"`
	V16               float32   `xorm:"FLOAT(8)"`
	V17               float32   `xorm:"FLOAT(8)"`
	V18               float32   `xorm:"FLOAT(8)"`
	V19               float32   `xorm:"FLOAT(8)"`
	V20               float32   `xorm:"FLOAT(8)"`
	V21               float32   `xorm:"FLOAT(8)"`
	V22               float32   `xorm:"FLOAT(8)"`
	V23               float32   `xorm:"FLOAT(8)"`
	V24               float32   `xorm:"FLOAT(8)"`
	V25               float32   `xorm:"FLOAT(8)"`
	V26               float32   `xorm:"FLOAT(8)"`
	V27               float32   `xorm:"FLOAT(8)"`
	V28               float32   `xorm:"FLOAT(8)"`
	V29               float32   `xorm:"FLOAT(8)"`
	V30               float32   `xorm:"FLOAT(8)"`
	V31               float32   `xorm:"FLOAT(8)"`
	V32               float32   `xorm:"FLOAT(8)"`
	V33               float32   `xorm:"FLOAT(8)"`
	V34               float32   `xorm:"FLOAT(8)"`
	V35               float32   `xorm:"FLOAT(8)"`
	V36               float32   `xorm:"FLOAT(8)"`
	V37               float32   `xorm:"FLOAT(8)"`
	V38               float32   `xorm:"FLOAT(8)"`
	V39               float32   `xorm:"FLOAT(8)"`
	V40               float32   `xorm:"FLOAT(8)"`
	V41               float32   `xorm:"FLOAT(8)"`
	V42               float32   `xorm:"FLOAT(8)"`
	V43               float32   `xorm:"FLOAT(8)"`
	V44               float32   `xorm:"FLOAT(8)"`
	V45               float32   `xorm:"FLOAT(8)"`
	V46               float32   `xorm:"FLOAT(8)"`
	V47               float32   `xorm:"FLOAT(8)"`
	V48               float32   `xorm:"FLOAT(8)"`
	V49               float32   `xorm:"FLOAT(8)"`
	V50               float32   `xorm:"FLOAT(8)"`
	V51               float32   `xorm:"FLOAT(8)"`
	V52               float32   `xorm:"FLOAT(8)"`
	V53               float32   `xorm:"FLOAT(8)"`
	V54               float32   `xorm:"FLOAT(8)"`
	V55               float32   `xorm:"FLOAT(8)"`
	V56               float32   `xorm:"FLOAT(8)"`
	V57               float32   `xorm:"FLOAT(8)"`
	V58               float32   `xorm:"FLOAT(8)"`
	V59               float32   `xorm:"FLOAT(8)"`
	V60               float32   `xorm:"FLOAT(8)"`
	V61               float32   `xorm:"FLOAT(8)"`
	V62               float32   `xorm:"FLOAT(8)"`
	V63               float32   `xorm:"FLOAT(8)"`
	V64               float32   `xorm:"FLOAT(8)"`
	V65               float32   `xorm:"FLOAT(8)"`
	V66               float32   `xorm:"FLOAT(8)"`
	V67               float32   `xorm:"FLOAT(8)"`
	V68               float32   `xorm:"FLOAT(8)"`
	V69               float32   `xorm:"FLOAT(8)"`
	V70               float32   `xorm:"FLOAT(8)"`
	V71               float32   `xorm:"FLOAT(8)"`
	V72               float32   `xorm:"FLOAT(8)"`
	V73               float32   `xorm:"FLOAT(8)"`
	V74               float32   `xorm:"FLOAT(8)"`
	V75               float32   `xorm:"FLOAT(8)"`
	V76               float32   `xorm:"FLOAT(8)"`
	V77               float32   `xorm:"FLOAT(8)"`
	V78               float32   `xorm:"FLOAT(8)"`
	V79               float32   `xorm:"FLOAT(8)"`
	V80               float32   `xorm:"FLOAT(8)"`
	V81               float32   `xorm:"FLOAT(8)"`
	V82               float32   `xorm:"FLOAT(8)"`
	V83               float32   `xorm:"FLOAT(8)"`
	V84               float32   `xorm:"FLOAT(8)"`
	V85               float32   `xorm:"FLOAT(8)"`
	V86               float32   `xorm:"FLOAT(8)"`
	V87               float32   `xorm:"FLOAT(8)"`
	V88               float32   `xorm:"FLOAT(8)"`
	V89               float32   `xorm:"FLOAT(8)"`
	V90               float32   `xorm:"FLOAT(8)"`
	V91               float32   `xorm:"FLOAT(8)"`
	V92               float32   `xorm:"FLOAT(8)"`
	V93               float32   `xorm:"FLOAT(8)"`
	V94               float32   `xorm:"FLOAT(8)"`
	V95               float32   `xorm:"FLOAT(8)"`
	H0                float32   `xorm:"FLOAT(8)"`
	H1                float32   `xorm:"FLOAT(8)"`
	H2                float32   `xorm:"FLOAT(8)"`
	H3                float32   `xorm:"FLOAT(8)"`
	H4                float32   `xorm:"FLOAT(8)"`
	H5                float32   `xorm:"FLOAT(8)"`
	H6                float32   `xorm:"FLOAT(8)"`
	H7                float32   `xorm:"FLOAT(8)"`
	H8                float32   `xorm:"FLOAT(8)"`
	H9                float32   `xorm:"FLOAT(8)"`
	H10               float32   `xorm:"FLOAT(8)"`
	H11               float32   `xorm:"FLOAT(8)"`
	H12               float32   `xorm:"FLOAT(8)"`
	H13               float32   `xorm:"FLOAT(8)"`
	H14               float32   `xorm:"FLOAT(8)"`
	H15               float32   `xorm:"FLOAT(8)"`
	H16               float32   `xorm:"FLOAT(8)"`
	H17               float32   `xorm:"FLOAT(8)"`
	H18               float32   `xorm:"FLOAT(8)"`
	H19               float32   `xorm:"FLOAT(8)"`
	H20               float32   `xorm:"FLOAT(8)"`
	H21               float32   `xorm:"FLOAT(8)"`
	H22               float32   `xorm:"FLOAT(8)"`
	H23               float32   `xorm:"FLOAT(8)"`
	H24               float32   `xorm:"FLOAT(8)"`
	H25               float32   `xorm:"FLOAT(8)"`
	H26               float32   `xorm:"FLOAT(8)"`
	H27               float32   `xorm:"FLOAT(8)"`
	H28               float32   `xorm:"FLOAT(8)"`
	H29               float32   `xorm:"FLOAT(8)"`
	H30               float32   `xorm:"FLOAT(8)"`
	H31               float32   `xorm:"FLOAT(8)"`
	H32               float32   `xorm:"FLOAT(8)"`
	H33               float32   `xorm:"FLOAT(8)"`
	H34               float32   `xorm:"FLOAT(8)"`
	H35               float32   `xorm:"FLOAT(8)"`
	H36               float32   `xorm:"FLOAT(8)"`
	H37               float32   `xorm:"FLOAT(8)"`
	H38               float32   `xorm:"FLOAT(8)"`
	H39               float32   `xorm:"FLOAT(8)"`
	H40               float32   `xorm:"FLOAT(8)"`
	H41               float32   `xorm:"FLOAT(8)"`
	H42               float32   `xorm:"FLOAT(8)"`
	H43               float32   `xorm:"FLOAT(8)"`
	H44               float32   `xorm:"FLOAT(8)"`
	H45               float32   `xorm:"FLOAT(8)"`
	H46               float32   `xorm:"FLOAT(8)"`
	H47               float32   `xorm:"FLOAT(8)"`
	H48               float32   `xorm:"FLOAT(8)"`
	H49               float32   `xorm:"FLOAT(8)"`
	H50               float32   `xorm:"FLOAT(8)"`
	H51               float32   `xorm:"FLOAT(8)"`
	H52               float32   `xorm:"FLOAT(8)"`
	H53               float32   `xorm:"FLOAT(8)"`
	H54               float32   `xorm:"FLOAT(8)"`
	H55               float32   `xorm:"FLOAT(8)"`
	H56               float32   `xorm:"FLOAT(8)"`
	H57               float32   `xorm:"FLOAT(8)"`
	H58               float32   `xorm:"FLOAT(8)"`
	H59               float32   `xorm:"FLOAT(8)"`
	H60               float32   `xorm:"FLOAT(8)"`
	H61               float32   `xorm:"FLOAT(8)"`
	H62               float32   `xorm:"FLOAT(8)"`
	H63               float32   `xorm:"FLOAT(8)"`
	H64               float32   `xorm:"FLOAT(8)"`
	H65               float32   `xorm:"FLOAT(8)"`
	H66               float32   `xorm:"FLOAT(8)"`
	H67               float32   `xorm:"FLOAT(8)"`
	H68               float32   `xorm:"FLOAT(8)"`
	H69               float32   `xorm:"FLOAT(8)"`
	H70               float32   `xorm:"FLOAT(8)"`
	H71               float32   `xorm:"FLOAT(8)"`
	H72               float32   `xorm:"FLOAT(8)"`
	H73               float32   `xorm:"FLOAT(8)"`
	H74               float32   `xorm:"FLOAT(8)"`
	H75               float32   `xorm:"FLOAT(8)"`
	H76               float32   `xorm:"FLOAT(8)"`
	H77               float32   `xorm:"FLOAT(8)"`
	H78               float32   `xorm:"FLOAT(8)"`
	H79               float32   `xorm:"FLOAT(8)"`
	H80               float32   `xorm:"FLOAT(8)"`
	H81               float32   `xorm:"FLOAT(8)"`
	H82               float32   `xorm:"FLOAT(8)"`
	H83               float32   `xorm:"FLOAT(8)"`
	H84               float32   `xorm:"FLOAT(8)"`
	H85               float32   `xorm:"FLOAT(8)"`
	H86               float32   `xorm:"FLOAT(8)"`
	H87               float32   `xorm:"FLOAT(8)"`
	H88               float32   `xorm:"FLOAT(8)"`
	H89               float32   `xorm:"FLOAT(8)"`
	H90               float32   `xorm:"FLOAT(8)"`
	H91               float32   `xorm:"FLOAT(8)"`
	H92               float32   `xorm:"FLOAT(8)"`
	H93               float32   `xorm:"FLOAT(8)"`
	H94               float32   `xorm:"FLOAT(8)"`
	H95               float32   `xorm:"FLOAT(8)"`
	S0                float32   `xorm:"FLOAT(8)"`
	S1                float32   `xorm:"FLOAT(8)"`
	S2                float32   `xorm:"FLOAT(8)"`
	S3                float32   `xorm:"FLOAT(8)"`
	S4                float32   `xorm:"FLOAT(8)"`
	S5                float32   `xorm:"FLOAT(8)"`
	S6                float32   `xorm:"FLOAT(8)"`
	S7                float32   `xorm:"FLOAT(8)"`
	S8                float32   `xorm:"FLOAT(8)"`
	S9                float32   `xorm:"FLOAT(8)"`
	S10               float32   `xorm:"FLOAT(8)"`
	S11               float32   `xorm:"FLOAT(8)"`
	S12               float32   `xorm:"FLOAT(8)"`
	S13               float32   `xorm:"FLOAT(8)"`
	S14               float32   `xorm:"FLOAT(8)"`
	S15               float32   `xorm:"FLOAT(8)"`
	S16               float32   `xorm:"FLOAT(8)"`
	S17               float32   `xorm:"FLOAT(8)"`
	S18               float32   `xorm:"FLOAT(8)"`
	S19               float32   `xorm:"FLOAT(8)"`
	S20               float32   `xorm:"FLOAT(8)"`
	S21               float32   `xorm:"FLOAT(8)"`
	S22               float32   `xorm:"FLOAT(8)"`
	S23               float32   `xorm:"FLOAT(8)"`
	S24               float32   `xorm:"FLOAT(8)"`
	S25               float32   `xorm:"FLOAT(8)"`
	S26               float32   `xorm:"FLOAT(8)"`
	S27               float32   `xorm:"FLOAT(8)"`
	S28               float32   `xorm:"FLOAT(8)"`
	S29               float32   `xorm:"FLOAT(8)"`
	S30               float32   `xorm:"FLOAT(8)"`
	S31               float32   `xorm:"FLOAT(8)"`
	S32               float32   `xorm:"FLOAT(8)"`
	S33               float32   `xorm:"FLOAT(8)"`
	S34               float32   `xorm:"FLOAT(8)"`
	S35               float32   `xorm:"FLOAT(8)"`
	S36               float32   `xorm:"FLOAT(8)"`
	S37               float32   `xorm:"FLOAT(8)"`
	S38               float32   `xorm:"FLOAT(8)"`
	S39               float32   `xorm:"FLOAT(8)"`
	S40               float32   `xorm:"FLOAT(8)"`
	S41               float32   `xorm:"FLOAT(8)"`
	S42               float32   `xorm:"FLOAT(8)"`
	S43               float32   `xorm:"FLOAT(8)"`
	S44               float32   `xorm:"FLOAT(8)"`
	S45               float32   `xorm:"FLOAT(8)"`
	S46               float32   `xorm:"FLOAT(8)"`
	S47               float32   `xorm:"FLOAT(8)"`
	S48               float32   `xorm:"FLOAT(8)"`
	S49               float32   `xorm:"FLOAT(8)"`
	S50               float32   `xorm:"FLOAT(8)"`
	S51               float32   `xorm:"FLOAT(8)"`
	S52               float32   `xorm:"FLOAT(8)"`
	S53               float32   `xorm:"FLOAT(8)"`
	S54               float32   `xorm:"FLOAT(8)"`
	S55               float32   `xorm:"FLOAT(8)"`
	S56               float32   `xorm:"FLOAT(8)"`
	S57               float32   `xorm:"FLOAT(8)"`
	S58               float32   `xorm:"FLOAT(8)"`
	S59               float32   `xorm:"FLOAT(8)"`
	S60               float32   `xorm:"FLOAT(8)"`
	S61               float32   `xorm:"FLOAT(8)"`
	S62               float32   `xorm:"FLOAT(8)"`
	S63               float32   `xorm:"FLOAT(8)"`
	S64               float32   `xorm:"FLOAT(8)"`
	S65               float32   `xorm:"FLOAT(8)"`
	S66               float32   `xorm:"FLOAT(8)"`
	S67               float32   `xorm:"FLOAT(8)"`
	S68               float32   `xorm:"FLOAT(8)"`
	S69               float32   `xorm:"FLOAT(8)"`
	S70               float32   `xorm:"FLOAT(8)"`
	S71               float32   `xorm:"FLOAT(8)"`
	S72               float32   `xorm:"FLOAT(8)"`
	S73               float32   `xorm:"FLOAT(8)"`
	S74               float32   `xorm:"FLOAT(8)"`
	S75               float32   `xorm:"FLOAT(8)"`
	S76               float32   `xorm:"FLOAT(8)"`
	S77               float32   `xorm:"FLOAT(8)"`
	S78               float32   `xorm:"FLOAT(8)"`
	S79               float32   `xorm:"FLOAT(8)"`
	S80               float32   `xorm:"FLOAT(8)"`
	S81               float32   `xorm:"FLOAT(8)"`
	S82               float32   `xorm:"FLOAT(8)"`
	S83               float32   `xorm:"FLOAT(8)"`
	S84               float32   `xorm:"FLOAT(8)"`
	S85               float32   `xorm:"FLOAT(8)"`
	S86               float32   `xorm:"FLOAT(8)"`
	S87               float32   `xorm:"FLOAT(8)"`
	S88               float32   `xorm:"FLOAT(8)"`
	S89               float32   `xorm:"FLOAT(8)"`
	S90               float32   `xorm:"FLOAT(8)"`
	S91               float32   `xorm:"FLOAT(8)"`
	S92               float32   `xorm:"FLOAT(8)"`
	S93               float32   `xorm:"FLOAT(8)"`
	S94               float32   `xorm:"FLOAT(8)"`
	S95               float32   `xorm:"FLOAT(8)"`
	A0                float32   `xorm:"FLOAT(8)"`
	A1                float32   `xorm:"FLOAT(8)"`
	A2                float32   `xorm:"FLOAT(8)"`
	A3                float32   `xorm:"FLOAT(8)"`
	A4                float32   `xorm:"FLOAT(8)"`
	A5                float32   `xorm:"FLOAT(8)"`
	A6                float32   `xorm:"FLOAT(8)"`
	A7                float32   `xorm:"FLOAT(8)"`
	A8                float32   `xorm:"FLOAT(8)"`
	A9                float32   `xorm:"FLOAT(8)"`
	A10               float32   `xorm:"FLOAT(8)"`
	A11               float32   `xorm:"FLOAT(8)"`
	A12               float32   `xorm:"FLOAT(8)"`
	A13               float32   `xorm:"FLOAT(8)"`
	A14               float32   `xorm:"FLOAT(8)"`
	A15               float32   `xorm:"FLOAT(8)"`
	A16               float32   `xorm:"FLOAT(8)"`
	A17               float32   `xorm:"FLOAT(8)"`
	A18               float32   `xorm:"FLOAT(8)"`
	A19               float32   `xorm:"FLOAT(8)"`
	A20               float32   `xorm:"FLOAT(8)"`
	A21               float32   `xorm:"FLOAT(8)"`
	A22               float32   `xorm:"FLOAT(8)"`
	A23               float32   `xorm:"FLOAT(8)"`
	A24               float32   `xorm:"FLOAT(8)"`
	A25               float32   `xorm:"FLOAT(8)"`
	A26               float32   `xorm:"FLOAT(8)"`
	A27               float32   `xorm:"FLOAT(8)"`
	A28               float32   `xorm:"FLOAT(8)"`
	A29               float32   `xorm:"FLOAT(8)"`
	A30               float32   `xorm:"FLOAT(8)"`
	A31               float32   `xorm:"FLOAT(8)"`
	A32               float32   `xorm:"FLOAT(8)"`
	A33               float32   `xorm:"FLOAT(8)"`
	A34               float32   `xorm:"FLOAT(8)"`
	A35               float32   `xorm:"FLOAT(8)"`
	A36               float32   `xorm:"FLOAT(8)"`
	A37               float32   `xorm:"FLOAT(8)"`
	A38               float32   `xorm:"FLOAT(8)"`
	A39               float32   `xorm:"FLOAT(8)"`
	A40               float32   `xorm:"FLOAT(8)"`
	A41               float32   `xorm:"FLOAT(8)"`
	A42               float32   `xorm:"FLOAT(8)"`
	A43               float32   `xorm:"FLOAT(8)"`
	A44               float32   `xorm:"FLOAT(8)"`
	A45               float32   `xorm:"FLOAT(8)"`
	A46               float32   `xorm:"FLOAT(8)"`
	A47               float32   `xorm:"FLOAT(8)"`
	A48               float32   `xorm:"FLOAT(8)"`
	A49               float32   `xorm:"FLOAT(8)"`
	A50               float32   `xorm:"FLOAT(8)"`
	A51               float32   `xorm:"FLOAT(8)"`
	A52               float32   `xorm:"FLOAT(8)"`
	A53               float32   `xorm:"FLOAT(8)"`
	A54               float32   `xorm:"FLOAT(8)"`
	A55               float32   `xorm:"FLOAT(8)"`
	A56               float32   `xorm:"FLOAT(8)"`
	A57               float32   `xorm:"FLOAT(8)"`
	A58               float32   `xorm:"FLOAT(8)"`
	A59               float32   `xorm:"FLOAT(8)"`
	A60               float32   `xorm:"FLOAT(8)"`
	A61               float32   `xorm:"FLOAT(8)"`
	A62               float32   `xorm:"FLOAT(8)"`
	A63               float32   `xorm:"FLOAT(8)"`
	A64               float32   `xorm:"FLOAT(8)"`
	A65               float32   `xorm:"FLOAT(8)"`
	A66               float32   `xorm:"FLOAT(8)"`
	A67               float32   `xorm:"FLOAT(8)"`
	A68               float32   `xorm:"FLOAT(8)"`
	A69               float32   `xorm:"FLOAT(8)"`
	A70               float32   `xorm:"FLOAT(8)"`
	A71               float32   `xorm:"FLOAT(8)"`
	A72               float32   `xorm:"FLOAT(8)"`
	A73               float32   `xorm:"FLOAT(8)"`
	A74               float32   `xorm:"FLOAT(8)"`
	A75               float32   `xorm:"FLOAT(8)"`
	A76               float32   `xorm:"FLOAT(8)"`
	A77               float32   `xorm:"FLOAT(8)"`
	A78               float32   `xorm:"FLOAT(8)"`
	A79               float32   `xorm:"FLOAT(8)"`
	A80               float32   `xorm:"FLOAT(8)"`
	A81               float32   `xorm:"FLOAT(8)"`
	A82               float32   `xorm:"FLOAT(8)"`
	A83               float32   `xorm:"FLOAT(8)"`
	A84               float32   `xorm:"FLOAT(8)"`
	A85               float32   `xorm:"FLOAT(8)"`
	A86               float32   `xorm:"FLOAT(8)"`
	A87               float32   `xorm:"FLOAT(8)"`
	A88               float32   `xorm:"FLOAT(8)"`
	A89               float32   `xorm:"FLOAT(8)"`
	A90               float32   `xorm:"FLOAT(8)"`
	A91               float32   `xorm:"FLOAT(8)"`
	A92               float32   `xorm:"FLOAT(8)"`
	A93               float32   `xorm:"FLOAT(8)"`
	A94               float32   `xorm:"FLOAT(8)"`
	A95               float32   `xorm:"FLOAT(8)"`
	M0                float32   `xorm:"FLOAT(8)"`
	M1                float32   `xorm:"FLOAT(8)"`
	M2                float32   `xorm:"FLOAT(8)"`
	M3                float32   `xorm:"FLOAT(8)"`
	M4                float32   `xorm:"FLOAT(8)"`
	M5                float32   `xorm:"FLOAT(8)"`
	M6                float32   `xorm:"FLOAT(8)"`
	M7                float32   `xorm:"FLOAT(8)"`
	M8                float32   `xorm:"FLOAT(8)"`
	M9                float32   `xorm:"FLOAT(8)"`
	M10               float32   `xorm:"FLOAT(8)"`
	M11               float32   `xorm:"FLOAT(8)"`
	M12               float32   `xorm:"FLOAT(8)"`
	M13               float32   `xorm:"FLOAT(8)"`
	M14               float32   `xorm:"FLOAT(8)"`
	M15               float32   `xorm:"FLOAT(8)"`
	M16               float32   `xorm:"FLOAT(8)"`
	M17               float32   `xorm:"FLOAT(8)"`
	M18               float32   `xorm:"FLOAT(8)"`
	M19               float32   `xorm:"FLOAT(8)"`
	M20               float32   `xorm:"FLOAT(8)"`
	M21               float32   `xorm:"FLOAT(8)"`
	M22               float32   `xorm:"FLOAT(8)"`
	M23               float32   `xorm:"FLOAT(8)"`
	M24               float32   `xorm:"FLOAT(8)"`
	M25               float32   `xorm:"FLOAT(8)"`
	M26               float32   `xorm:"FLOAT(8)"`
	M27               float32   `xorm:"FLOAT(8)"`
	M28               float32   `xorm:"FLOAT(8)"`
	M29               float32   `xorm:"FLOAT(8)"`
	M30               float32   `xorm:"FLOAT(8)"`
	M31               float32   `xorm:"FLOAT(8)"`
	M32               float32   `xorm:"FLOAT(8)"`
	M33               float32   `xorm:"FLOAT(8)"`
	M34               float32   `xorm:"FLOAT(8)"`
	M35               float32   `xorm:"FLOAT(8)"`
	M36               float32   `xorm:"FLOAT(8)"`
	M37               float32   `xorm:"FLOAT(8)"`
	M38               float32   `xorm:"FLOAT(8)"`
	M39               float32   `xorm:"FLOAT(8)"`
	M40               float32   `xorm:"FLOAT(8)"`
	M41               float32   `xorm:"FLOAT(8)"`
	M42               float32   `xorm:"FLOAT(8)"`
	M43               float32   `xorm:"FLOAT(8)"`
	M44               float32   `xorm:"FLOAT(8)"`
	M45               float32   `xorm:"FLOAT(8)"`
	M46               float32   `xorm:"FLOAT(8)"`
	M47               float32   `xorm:"FLOAT(8)"`
	M48               float32   `xorm:"FLOAT(8)"`
	M49               float32   `xorm:"FLOAT(8)"`
	M50               float32   `xorm:"FLOAT(8)"`
	M51               float32   `xorm:"FLOAT(8)"`
	M52               float32   `xorm:"FLOAT(8)"`
	M53               float32   `xorm:"FLOAT(8)"`
	M54               float32   `xorm:"FLOAT(8)"`
	M55               float32   `xorm:"FLOAT(8)"`
	M56               float32   `xorm:"FLOAT(8)"`
	M57               float32   `xorm:"FLOAT(8)"`
	M58               float32   `xorm:"FLOAT(8)"`
	M59               float32   `xorm:"FLOAT(8)"`
	M60               float32   `xorm:"FLOAT(8)"`
	M61               float32   `xorm:"FLOAT(8)"`
	M62               float32   `xorm:"FLOAT(8)"`
	M63               float32   `xorm:"FLOAT(8)"`
	M64               float32   `xorm:"FLOAT(8)"`
	M65               float32   `xorm:"FLOAT(8)"`
	M66               float32   `xorm:"FLOAT(8)"`
	M67               float32   `xorm:"FLOAT(8)"`
	M68               float32   `xorm:"FLOAT(8)"`
	M69               float32   `xorm:"FLOAT(8)"`
	M70               float32   `xorm:"FLOAT(8)"`
	M71               float32   `xorm:"FLOAT(8)"`
	M72               float32   `xorm:"FLOAT(8)"`
	M73               float32   `xorm:"FLOAT(8)"`
	M74               float32   `xorm:"FLOAT(8)"`
	M75               float32   `xorm:"FLOAT(8)"`
	M76               float32   `xorm:"FLOAT(8)"`
	M77               float32   `xorm:"FLOAT(8)"`
	M78               float32   `xorm:"FLOAT(8)"`
	M79               float32   `xorm:"FLOAT(8)"`
	M80               float32   `xorm:"FLOAT(8)"`
	M81               float32   `xorm:"FLOAT(8)"`
	M82               float32   `xorm:"FLOAT(8)"`
	M83               float32   `xorm:"FLOAT(8)"`
	M84               float32   `xorm:"FLOAT(8)"`
	M85               float32   `xorm:"FLOAT(8)"`
	M86               float32   `xorm:"FLOAT(8)"`
	M87               float32   `xorm:"FLOAT(8)"`
	M88               float32   `xorm:"FLOAT(8)"`
	M89               float32   `xorm:"FLOAT(8)"`
	M90               float32   `xorm:"FLOAT(8)"`
	M91               float32   `xorm:"FLOAT(8)"`
	M92               float32   `xorm:"FLOAT(8)"`
	M93               float32   `xorm:"FLOAT(8)"`
	M94               float32   `xorm:"FLOAT(8)"`
	M95               float32   `xorm:"FLOAT(8)"`
}

type Contingent struct {
	Contingentid   []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name           string `xorm:"not null NVARCHAR(510)"`
	Isweekday      int    `xorm:"not null BIT(1)"`
	Weekvalues     string `xorm:"NVARCHAR(510)"`
	Monthdayvalues string `xorm:"NVARCHAR(510)"`
}

type Forecastservicelevelgoal struct {
	Forecastservicelevelgoalid int       `xorm:"not null pk INT(4)"`
	Forecastdate               time.Time `xorm:"not null DATETIME(8)"`
	Forecastid                 int       `xorm:"not null INT(4)"`
	Sl0                        int       `xorm:"not null INT(4)"`
	Sl1                        int       `xorm:"not null INT(4)"`
	Sl2                        int       `xorm:"not null INT(4)"`
	Sl3                        int       `xorm:"not null INT(4)"`
	Sl4                        int       `xorm:"not null INT(4)"`
	Sl5                        int       `xorm:"not null INT(4)"`
	Sl6                        int       `xorm:"not null INT(4)"`
	Sl7                        int       `xorm:"not null INT(4)"`
	Sl8                        int       `xorm:"not null INT(4)"`
	Sl9                        int       `xorm:"not null INT(4)"`
	Sl10                       int       `xorm:"not null INT(4)"`
	Sl11                       int       `xorm:"not null INT(4)"`
	Sl12                       int       `xorm:"not null INT(4)"`
	Sl13                       int       `xorm:"not null INT(4)"`
	Sl14                       int       `xorm:"not null INT(4)"`
	Sl15                       int       `xorm:"not null INT(4)"`
	Sl16                       int       `xorm:"not null INT(4)"`
	Sl17                       int       `xorm:"not null INT(4)"`
	Sl18                       int       `xorm:"not null INT(4)"`
	Sl19                       int       `xorm:"not null INT(4)"`
	Sl20                       int       `xorm:"not null INT(4)"`
	Sl21                       int       `xorm:"not null INT(4)"`
	Sl22                       int       `xorm:"not null INT(4)"`
	Sl23                       int       `xorm:"not null INT(4)"`
}

type Campaign struct {
	Campaignid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name       string `xorm:"not null NVARCHAR(510)"`
}

type Multiswapshift struct {
	Multiswapshiftid    []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Employeeaid         []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Employeebid         []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Start               time.Time `xorm:"not null DATETIME(8)"`
	End                 time.Time `xorm:"not null DATETIME(8)"`
	Swapmanager         []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Swaptime            time.Time `xorm:"not null DATETIME(8)"`
	Warningmessagea     string    `xorm:"NVARCHAR(-1)"`
	Warningmessageb     string    `xorm:"NVARCHAR(-1)"`
	Substitutiongroupid []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Employeeainfo       string    `xorm:"not null NVARCHAR(256)"`
	Employeebinfo       string    `xorm:"not null NVARCHAR(256)"`
	Swapmanagerinfo     string    `xorm:"not null NVARCHAR(256)"`
}

type Forecasthistory struct {
	Forecasthistoryid int       `xorm:"not null pk INT(4)"`
	Forecastid        int       `xorm:"not null INT(4)"`
	Startat           time.Time `xorm:"not null DATETIME(8)"`
	Endat             time.Time `xorm:"not null DATETIME(8)"`
	Weight            int       `xorm:"not null default 10 INT(4)"`
	Color             string    `xorm:"not null NVARCHAR(20)"`
	Remark            string    `xorm:"NVARCHAR(200)"`
	Forecaststartdate time.Time `xorm:"not null DATETIME(8)"`
	Forecastenddate   time.Time `xorm:"not null DATETIME(8)"`
	Mappingpattern    int       `xorm:"not null INT(4)"`
}

type Laborrule struct {
	Laborruleid                    []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name                           string `xorm:"NVARCHAR(510)"`
	Maxovertime                    int64  `xorm:"BIGINT(8)"`
	Maxshrinked                    int64  `xorm:"BIGINT(8)"`
	Mcdo                           int    `xorm:"INT(4)"`
	Mcwd                           int    `xorm:"INT(4)"`
	Minidlegap                     int64  `xorm:"BIGINT(8)"`
	Stddailylaborhour              int64  `xorm:"BIGINT(8)"`
	Maxlaborhour                   int64  `xorm:"BIGINT(8)"`
	Minlaborhour                   int64  `xorm:"BIGINT(8)"`
	Maxswaptimes                   int    `xorm:"INT(4)"`
	Add1dayoffeachsaturdayincs     int    `xorm:"BIT(1)"`
	Add1dayoffeachsundayincs       int    `xorm:"BIT(1)"`
	Add1dayoffeachholidayincs      int    `xorm:"BIT(1)"`
	Holidayshiftrule               int    `xorm:"INT(4)"`
	Maxfwtimes                     int    `xorm:"INT(4)"`
	Minfwtimes                     int    `xorm:"INT(4)"`
	Maxpwtimes                     int    `xorm:"INT(4)"`
	Minpwtimes                     int    `xorm:"INT(4)"`
	Isgrouping                     int    `xorm:"BIT(1)"`
	Ismappingevent                 int    `xorm:"BIT(1)"`
	Specifiednumberofdays          int    `xorm:"INT(4)"`
	Systemaccumulate               int    `xorm:"BIT(1)"`
	Validdayoffweekdayfromsun2sat  int    `xorm:"INT(4)"`
	Validdayoffmonthday            int    `xorm:"INT(4)"`
	Color                          string `xorm:"NVARCHAR(510)"`
	Validworkingweekdayfromsun2sat int    `xorm:"INT(4)"`
	Tocwdnames                     string `xorm:"NVARCHAR(2000)"`
}

type Calendarevent struct {
	Id                int       `xorm:"not null pk INT(4)"`
	Calendareventtype string    `xorm:"not null NVARCHAR(510)"`
	Eventname         string    `xorm:"NVARCHAR(510)"`
	Startdate         time.Time `xorm:"not null DATETIME(8)"`
	Enddate           time.Time `xorm:"not null DATETIME(8)"`
	Country           string    `xorm:"NVARCHAR(510)"`
	Timezoneid        string    `xorm:"NVARCHAR(510)"`
}

type Schedulingpayload struct {
	Schedulingpayloadid           int `xorm:"not null pk INT(4)"`
	Isgrouping                    int `xorm:"BIT(1)"`
	Ismappingevent                int `xorm:"BIT(1)"`
	Validdayoffweekdayfromsun2sat int `xorm:"INT(4)"`
	Validdayoffmonthday           int `xorm:"INT(4)"`
}

type Biddingsetting struct {
	Biddingsettingid      []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Biddingannouncementid []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Biddingdate           time.Time `xorm:"not null DATETIME(8)"`
}

type Biddingannouncementorganization struct {
	Biddingannouncementid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Organizationid        []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Orgname               string `xorm:"NVARCHAR(256)"`
}

type Seat struct {
	Seatid              []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Areaid              []byte `xorm:"unique(UQ__tmp_ms_x__A7323A504979DDF4) UNIQUEIDENTIFIER(16)"`
	Number              string `xorm:"unique(UQ__tmp_ms_x__A7323A504979DDF4) NVARCHAR(20)"`
	Extno               string `xorm:"not null unique NVARCHAR(510)"`
	Inuse               int    `xorm:"not null BIT(1)"`
	Isopen              int    `xorm:"BIT(1)"`
	Rank                int    `xorm:"INT(4)"`
	Priority            int    `xorm:"INT(4)"`
	Locationindex       int    `xorm:"INT(4)"`
	Usingorganozationid []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Ip                  string `xorm:"NVARCHAR(510)"`
}

type Biddingannouncement struct {
	Biddingannouncementid []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name                  string    `xorm:"not null NVARCHAR(128)"`
	Start                 time.Time `xorm:"not null DATETIME(8)"`
	End                   time.Time `xorm:"not null DATETIME(8)"`
	Timeoffvalues         string    `xorm:"not null NVARCHAR(-1)"`
	Dayoffvalues          string    `xorm:"not null NVARCHAR(-1)"`
	Status                string    `xorm:"not null NVARCHAR(32)"`
	Maxdayoffcount        int       `xorm:"not null INT(4)"`
	Agentcanseecount      int       `xorm:"not null INT(4)"`
}

type Customemployeegroup struct {
	Customemployeegroupid int    `xorm:"not null pk INT(4)"`
	Groupname             string `xorm:"NVARCHAR(510)"`
	Owner                 []byte `xorm:"not null UNIQUEIDENTIFIER(16)"`
}

type Resource struct {
	Resourceid  []byte `xorm:"not null pk unique UNIQUEIDENTIFIER(16)"`
	Name        string `xorm:"NVARCHAR(510)"`
	Value       string `xorm:"NVARCHAR(510)"`
	Resourcekey string `xorm:"NVARCHAR(510)"`
	Groupname   string `xorm:"NVARCHAR(510)"`
	Isused      int    `xorm:"BIT(1)"`
	Maxvalue    string `xorm:"DECIMAL(18)"`
	Minvalue    string `xorm:"DECIMAL(18)"`
	Stepvalue   string `xorm:"DECIMAL(18)"`
	Childtype   string `xorm:"NVARCHAR(510)"`
	Valuetype   string `xorm:"NVARCHAR(510)"`
	Moduleid    []byte `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Priorityemployee struct {
	Priorityemployeeid int64  `xorm:"not null pk BIGINT(8)"`
	Seat               []byte `xorm:"unique(UQ__Priority__A585CEDA4E1E9780) UNIQUEIDENTIFIER(16)"`
	Employeeid         []byte `xorm:"unique(UQ__Priority__A585CEDA4E1E9780) UNIQUEIDENTIFIER(16)"`
	Priority           int    `xorm:"INT(4)"`
}

type Skill struct {
	Skillid     []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name        string `xorm:"NVARCHAR(510)"`
	Description string `xorm:"NVARCHAR(510)"`
}

type Shiftgroup struct {
	Shiftgroupid                  []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name                          string `xorm:"NVARCHAR(510)"`
	Isoneassignmenttypeonly       int    `xorm:"BIT(1)"`
	Maxgrouptimes                 int    `xorm:"INT(4)"`
	Mingrouptimes                 int    `xorm:"INT(4)"`
	Maxcontinuegrouptimes         int    `xorm:"INT(4)"`
	Mincontinuegrouptimes         int    `xorm:"INT(4)"`
	Mindayofftimespan             int    `xorm:"INT(4)"`
	Priority                      int    `xorm:"INT(4)"`
	Inuse                         int    `xorm:"BIT(1)"`
	Validdayoffweekdayfromsun2sat int    `xorm:"INT(4)"`
	Validdayoffmonthday           int    `xorm:"INT(4)"`
}

type Campaignschedule struct {
	Campaignscheduleid []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Campaign           []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Name               string    `xorm:"NVARCHAR(510)"`
	Starttime          time.Time `xorm:"DATETIME(8)"`
	Endtime            time.Time `xorm:"DATETIME(8)"`
	Progress           string    `xorm:"NVARCHAR(32)"`
	Seatcapacity       int       `xorm:"INT(4)"`
	Maxseat            int       `xorm:"INT(4)"`
	Lockestimates      int       `xorm:"BIT(1)"`
	Opendate           time.Time `xorm:"DATETIME(8)"`
	Lastopened         time.Time `xorm:"DATETIME(8)"`
}

type Site struct {
	Siteid      []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name        string `xorm:"NVARCHAR(510)"`
	Description string `xorm:"NVARCHAR(510)"`
}

type Organizationseatingarea struct {
	Organizationseatingareaid int    `xorm:"not null pk INT(4)"`
	Areaid                    []byte `xorm:"unique(UQ__Organiza__4C1590F855BFB948) UNIQUEIDENTIFIER(16)"`
	Organizationid            []byte `xorm:"unique(UQ__Organiza__4C1590F855BFB948) UNIQUEIDENTIFIER(16)"`
	Priority                  int    `xorm:"INT(4)"`
	Isselected                int    `xorm:"BIT(1)"`
	Targetseat                []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Methodology               string `xorm:"NVARCHAR(20)"`
}

type Substitutiongroup struct {
	Substitutiongroupid             []byte  `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name                            string  `xorm:"not null NVARCHAR(64)"`
	Skillproficiencydifferencerange float32 `xorm:"not null FLOAT(8)"`
	Swapexpirebeforedays            int     `xorm:"not null INT(4)"`
	Swapseexpirebeforedays          int     `xorm:"not null INT(4)"`
	Swapstartatintervalallowance    int     `xorm:"not null INT(4)"`
	Semaxspanlength                 int     `xorm:"not null INT(4)"`
	Isswapautodeny                  int     `xorm:"not null BIT(1)"`
	Isswapautoapprove               int     `xorm:"not null BIT(1)"`
	Canswapsubevent                 int     `xorm:"not null BIT(1)"`
	Iscountofsubstituteforapplier   int     `xorm:"not null BIT(1)"`
	Iscountofsubstituteforreplier   int     `xorm:"not null BIT(1)"`
	Iscountofswapforapplier         int     `xorm:"not null BIT(1)"`
	Iscountofswapforreplier         int     `xorm:"not null BIT(1)"`
	Iscountofmultiswapforapplier    int     `xorm:"not null BIT(1)"`
	Iscountofmultiswapforreplier    int     `xorm:"not null BIT(1)"`
	Notificationtype                string  `xorm:"not null NVARCHAR(32)"`
	Crossingnumberforl1             int     `xorm:"not null INT(4)"`
	Limitforl1                      int     `xorm:"not null BIT(1)"`
	Crossingnumberforl2             int     `xorm:"not null INT(4)"`
	Limitforl2                      int     `xorm:"not null BIT(1)"`
	Maxenrollmentdateyears          float32 `xorm:"not null FLOAT(8)"`
	Mintimeslicelen                 int     `xorm:"not null default 15 INT(4)"`
	Unlaboredsubeventid             []byte  `xorm:"UNIQUEIDENTIFIER(16)"`
	Regularsubeventid               []byte  `xorm:"UNIQUEIDENTIFIER(16)"`
	Rookielearningperiod            int     `xorm:"not null default 1 INT(4)"`
	Canswapmultishift               int     `xorm:"not null default 0 BIT(1)"`
	Canswaptimeslice                int     `xorm:"not null default 0 BIT(1)"`
	Dailyworkinghours               int     `xorm:"INT(4)"`
	Initialtotalworkinghours        int     `xorm:"INT(4)"`
}

type Area struct {
	Areaid   []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name     string `xorm:"not null NVARCHAR(510)"`
	Priority int    `xorm:"INT(4)"`
	Rows     int    `xorm:"INT(4)"`
	Columns  int    `xorm:"INT(4)"`
	Siteid   []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Campaign []byte `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Shrinkagetype struct {
	Id                     int    `xorm:"not null pk INT(4)"`
	Name                   string `xorm:"not null NVARCHAR(50)"`
	Isonline               int    `xorm:"not null BIT(1)"`
	Alertlevel             int    `xorm:"not null INT(4)"`
	Isshrinkage            int    `xorm:"not null BIT(1)"`
	Minsecond              int    `xorm:"not null INT(4)"`
	Maxsecond              int    `xorm:"not null INT(4)"`
	Isstartonshiftstart    int    `xorm:"not null BIT(1)"`
	Isendonshiftstart      int    `xorm:"not null BIT(1)"`
	Isstartonshiftend      int    `xorm:"not null BIT(1)"`
	Isendonshiftend        int    `xorm:"not null BIT(1)"`
	Inused                 int    `xorm:"not null default 1 BIT(1)"`
	Isstartonsubeventstart int    `xorm:"not null default 0 BIT(1)"`
	Isendonsubeventstart   int    `xorm:"not null default 0 BIT(1)"`
	Isstartonsubeventend   int    `xorm:"not null default 0 BIT(1)"`
	Isendonsubeventend     int    `xorm:"not null default 0 BIT(1)"`
	Subeventids            string `xorm:"NVARCHAR(8000)"`
	Isstartonshiftrange    int    `xorm:"not null default 0 BIT(1)"`
	Isendonshiftrange      int    `xorm:"not null default 0 BIT(1)"`
	Isstartoutshiftrange   int    `xorm:"not null default 0 BIT(1)"`
	Isendoutshiftrange     int    `xorm:"not null default 0 BIT(1)"`
	Isstartonsubeventrange int    `xorm:"not null default 0 BIT(1)"`
	Isendonsubeventrange   int    `xorm:"not null default 0 BIT(1)"`
	Iscutting              int    `xorm:"not null default 0 BIT(1)"`
}

type Organization struct {
	Organizationid  []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name            string `xorm:"NVARCHAR(510)"`
	Parentid        []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Laborrule       []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Certificationid string `xorm:"NVARCHAR(80)"`
}

type Servicelevelgoal struct {
	Servicelevelgoalid int    `xorm:"not null pk INT(4)"`
	Name               string `xorm:"not null NVARCHAR(60)"`
	Description        string `xorm:"NVARCHAR(200)"`
	Servicesecond      int    `xorm:"not null INT(4)"`
	Servicelevel       int    `xorm:"INT(4)"`
	Goalpattern        int    `xorm:"not null INT(4)"`
}

type Eventtype struct {
	Id          int    `xorm:"not null pk INT(4)"`
	Name        string `xorm:"not null NVARCHAR(100)"`
	Eventbookid int    `xorm:"not null INT(4)"`
	Isdeleted   int    `xorm:"not null default 0 BIT(1)"`
}

type Laborruleshiftgroup struct {
	Laborruleid  []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Shiftgroupid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Dayservicelevelgoal struct {
	Dayservicelevelgoalid int `xorm:"not null pk INT(4)"`
	Servicelevelgoalid    int `xorm:"not null INT(4)"`
	Weekday               int `xorm:"not null INT(4)"`
	Sl0                   int `xorm:"not null INT(4)"`
	Sl1                   int `xorm:"not null INT(4)"`
	Sl2                   int `xorm:"not null INT(4)"`
	Sl3                   int `xorm:"not null INT(4)"`
	Sl4                   int `xorm:"not null INT(4)"`
	Sl5                   int `xorm:"not null INT(4)"`
	Sl6                   int `xorm:"not null INT(4)"`
	Sl7                   int `xorm:"not null INT(4)"`
	Sl8                   int `xorm:"not null INT(4)"`
	Sl9                   int `xorm:"not null INT(4)"`
	Sl10                  int `xorm:"not null INT(4)"`
	Sl11                  int `xorm:"not null INT(4)"`
	Sl12                  int `xorm:"not null INT(4)"`
	Sl13                  int `xorm:"not null INT(4)"`
	Sl14                  int `xorm:"not null INT(4)"`
	Sl15                  int `xorm:"not null INT(4)"`
	Sl16                  int `xorm:"not null INT(4)"`
	Sl17                  int `xorm:"not null INT(4)"`
	Sl18                  int `xorm:"not null INT(4)"`
	Sl19                  int `xorm:"not null INT(4)"`
	Sl20                  int `xorm:"not null INT(4)"`
	Sl21                  int `xorm:"not null INT(4)"`
	Sl22                  int `xorm:"not null INT(4)"`
	Sl23                  int `xorm:"not null INT(4)"`
}

type Seatconsolidationrule struct {
	Seatconsolidationruleid int64  `xorm:"not null pk BIGINT(8)"`
	Title                   string `xorm:"NVARCHAR(510)"`
	Startvalue              int    `xorm:"INT(4)"`
	Endvalue                int    `xorm:"INT(4)"`
	Priority                int    `xorm:"INT(4)"`
	Targetseat              []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Site                    []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Maxrank                 int    `xorm:"INT(4)"`
	Minrank                 int    `xorm:"INT(4)"`
	Matchwholeskills        int    `xorm:"BIT(1)"`
	Methodology             string `xorm:"NVARCHAR(20)"`
	Su                      int    `xorm:"BIT(1)"`
	Mo                      int    `xorm:"BIT(1)"`
	Tu                      int    `xorm:"BIT(1)"`
	We                      int    `xorm:"BIT(1)"`
	Th                      int    `xorm:"BIT(1)"`
	Fr                      int    `xorm:"BIT(1)"`
	Sa                      int    `xorm:"BIT(1)"`
}

type Eventphrase struct {
	Id          int    `xorm:"not null pk INT(4)"`
	Eventbookid int    `xorm:"not null INT(4)"`
	Text        string `xorm:"not null NVARCHAR(400)"`
	Isdeleted   int    `xorm:"not null default 0 BIT(1)"`
}

type Servicequeue struct {
	Servicequeueid     []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Description        string `xorm:"NVARCHAR(510)"`
	Abandonrate        int    `xorm:"INT(4)"`
	Campaign           []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Name               string `xorm:"not null NVARCHAR(510)"`
	Mappedskill        []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Minimumonlines     int    `xorm:"INT(4)"`
	Servicelevelgoalid int    `xorm:"INT(4)"`
	Enable             int    `xorm:"not null default 1 BIT(1)"`
}

type Eventbook struct {
	Id        int    `xorm:"not null pk INT(4)"`
	Name      string `xorm:"not null NVARCHAR(100)"`
	Role      string `xorm:"NVARCHAR(100)"`
	Isdeleted int    `xorm:"not null default 0 BIT(1)"`
	Isscored  int    `xorm:"not null default 0 BIT(1)"`
}

type Accumulateholiday struct {
	Id              int64     `xorm:"pk BIGINT(8)"`
	Producttype     string    `xorm:"not null NVARCHAR(510)"`
	Employeeid      []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Date            time.Time `xorm:"DATETIME(8)"`
	Workinghour     int64     `xorm:"BIGINT(8)"`
	Termdisplaytext string    `xorm:"NVARCHAR(510)"`
}

type Shiftgroupassignmenttype struct {
	Shiftgroupid     []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Assignmenttypeid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Agentstatus struct {
	Agentstatusid       int64     `xorm:"not null pk BIGINT(8)"`
	Splitid             string    `xorm:"NVARCHAR(40)"`
	Agentacdid          string    `xorm:"index NVARCHAR(32)"`
	Extno               string    `xorm:"index NVARCHAR(100)"`
	Timestamp           time.Time `xorm:"index DATETIME(8)"`
	Agentstatustypecode string    `xorm:"not null NVARCHAR(510)"`
}

type Authrolefunction struct {
	Rolefunctionid int    `xorm:"not null pk INT(4)"`
	Keyname        string `xorm:"NVARCHAR(510)"`
	Roleid         []byte `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Seatbox struct {
	Seatboxid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Version   int64  `xorm:"not null BIGINT(8)"`
}

type Employeeskill struct {
	Employeeid   []byte  `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Productivity float32 `xorm:"FLOAT(8)"`
	Skillid      []byte  `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Agentacdid   int     `xorm:"not null pk INT(4)"`
	Color        string  `xorm:"NVARCHAR(510)"`
}

type Hrshiftcodemapping struct {
	Csrstermname string `xorm:"not null NVARCHAR(100)"`
	Hrshiftcode  string `xorm:"not null NVARCHAR(100)"`
}

type Employeetermstyle struct {
	Employeeid  []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Termstyleid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type HibernateUniqueKey struct {
	NextHi int64 `xorm:"BIGINT(8)"`
}

type Subeventinsertrules struct {
	Parentid       []byte `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Startvalue     int    `xorm:"not null INT(4)"`
	Endvalue       int    `xorm:"not null INT(4)"`
	Subevent       []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Subeventlength int    `xorm:"INT(4)"`
	Occurscale     int    `xorm:"INT(4)"`
}

type Deletedterm struct {
	Id         int64     `xorm:"pk BIGINT(8)"`
	Starttime  time.Time `xorm:"not null DATETIME(8)"`
	Finishtime time.Time `xorm:"not null DATETIME(8)"`
	Locked     int       `xorm:"not null BIT(1)"`
	Level      int       `xorm:"INT(4)"`
	Name       string    `xorm:"NVARCHAR(510)"`
	Timeboxid  []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Sqpayloadrecord struct {
	Id              int64  `xorm:"pk BIGINT(8)"`
	Payloadcategory string `xorm:"not null NVARCHAR(64)"`
	Tense           string `xorm:"not null NVARCHAR(20)"`
	V0              string `xorm:"not null NVARCHAR(512)"`
	V1              string `xorm:"not null NVARCHAR(512)"`
	V2              string `xorm:"not null NVARCHAR(512)"`
	V3              string `xorm:"not null NVARCHAR(512)"`
	V4              string `xorm:"not null NVARCHAR(512)"`
	V5              string `xorm:"not null NVARCHAR(512)"`
	V6              string `xorm:"not null NVARCHAR(512)"`
	V7              string `xorm:"not null NVARCHAR(512)"`
	V8              string `xorm:"not null NVARCHAR(512)"`
	V9              string `xorm:"not null NVARCHAR(512)"`
	V10             string `xorm:"not null NVARCHAR(512)"`
	V11             string `xorm:"not null NVARCHAR(512)"`
	V12             string `xorm:"not null NVARCHAR(512)"`
	V13             string `xorm:"not null NVARCHAR(512)"`
	V14             string `xorm:"not null NVARCHAR(512)"`
	V15             string `xorm:"not null NVARCHAR(512)"`
	V16             string `xorm:"not null NVARCHAR(512)"`
	V17             string `xorm:"not null NVARCHAR(512)"`
	V18             string `xorm:"not null NVARCHAR(512)"`
	V19             string `xorm:"not null NVARCHAR(512)"`
	V20             string `xorm:"not null NVARCHAR(512)"`
	V21             string `xorm:"not null NVARCHAR(512)"`
	V22             string `xorm:"not null NVARCHAR(512)"`
	V23             string `xorm:"not null NVARCHAR(512)"`
	V24             string `xorm:"not null NVARCHAR(512)"`
	V25             string `xorm:"not null NVARCHAR(512)"`
	V26             string `xorm:"not null NVARCHAR(512)"`
	V27             string `xorm:"not null NVARCHAR(512)"`
	V28             string `xorm:"not null NVARCHAR(512)"`
	V29             string `xorm:"not null NVARCHAR(512)"`
	V30             string `xorm:"not null NVARCHAR(512)"`
	V31             string `xorm:"not null NVARCHAR(512)"`
	V32             string `xorm:"not null NVARCHAR(512)"`
	V33             string `xorm:"not null NVARCHAR(512)"`
	V34             string `xorm:"not null NVARCHAR(512)"`
	V35             string `xorm:"not null NVARCHAR(512)"`
	V36             string `xorm:"not null NVARCHAR(512)"`
	V37             string `xorm:"not null NVARCHAR(512)"`
	V38             string `xorm:"not null NVARCHAR(512)"`
	V39             string `xorm:"not null NVARCHAR(512)"`
	V40             string `xorm:"not null NVARCHAR(512)"`
	V41             string `xorm:"not null NVARCHAR(512)"`
	V42             string `xorm:"not null NVARCHAR(512)"`
	V43             string `xorm:"not null NVARCHAR(512)"`
	V44             string `xorm:"not null NVARCHAR(512)"`
	V45             string `xorm:"not null NVARCHAR(512)"`
	V46             string `xorm:"not null NVARCHAR(512)"`
	V47             string `xorm:"not null NVARCHAR(512)"`
	V48             string `xorm:"not null NVARCHAR(512)"`
	V49             string `xorm:"not null NVARCHAR(512)"`
	V50             string `xorm:"not null NVARCHAR(512)"`
	V51             string `xorm:"not null NVARCHAR(512)"`
	V52             string `xorm:"not null NVARCHAR(512)"`
	V53             string `xorm:"not null NVARCHAR(512)"`
	V54             string `xorm:"not null NVARCHAR(512)"`
	V55             string `xorm:"not null NVARCHAR(512)"`
	V56             string `xorm:"not null NVARCHAR(512)"`
	V57             string `xorm:"not null NVARCHAR(512)"`
	V58             string `xorm:"not null NVARCHAR(512)"`
	V59             string `xorm:"not null NVARCHAR(512)"`
	V60             string `xorm:"not null NVARCHAR(512)"`
	V61             string `xorm:"not null NVARCHAR(512)"`
	V62             string `xorm:"not null NVARCHAR(512)"`
	V63             string `xorm:"not null NVARCHAR(512)"`
	V64             string `xorm:"not null NVARCHAR(512)"`
	V65             string `xorm:"not null NVARCHAR(512)"`
	V66             string `xorm:"not null NVARCHAR(512)"`
	V67             string `xorm:"not null NVARCHAR(512)"`
	V68             string `xorm:"not null NVARCHAR(512)"`
	V69             string `xorm:"not null NVARCHAR(512)"`
	V70             string `xorm:"not null NVARCHAR(512)"`
	V71             string `xorm:"not null NVARCHAR(512)"`
	V72             string `xorm:"not null NVARCHAR(512)"`
	V73             string `xorm:"not null NVARCHAR(512)"`
	V74             string `xorm:"not null NVARCHAR(512)"`
	V75             string `xorm:"not null NVARCHAR(512)"`
	V76             string `xorm:"not null NVARCHAR(512)"`
	V77             string `xorm:"not null NVARCHAR(512)"`
	V78             string `xorm:"not null NVARCHAR(512)"`
	V79             string `xorm:"not null NVARCHAR(512)"`
	V80             string `xorm:"not null NVARCHAR(512)"`
	V81             string `xorm:"not null NVARCHAR(512)"`
	V82             string `xorm:"not null NVARCHAR(512)"`
	V83             string `xorm:"not null NVARCHAR(512)"`
	V84             string `xorm:"not null NVARCHAR(512)"`
	V85             string `xorm:"not null NVARCHAR(512)"`
	V86             string `xorm:"not null NVARCHAR(512)"`
	V87             string `xorm:"not null NVARCHAR(512)"`
	V88             string `xorm:"not null NVARCHAR(512)"`
	V89             string `xorm:"not null NVARCHAR(512)"`
	V90             string `xorm:"not null NVARCHAR(512)"`
	V91             string `xorm:"not null NVARCHAR(512)"`
	V92             string `xorm:"not null NVARCHAR(512)"`
	V93             string `xorm:"not null NVARCHAR(512)"`
	V94             string `xorm:"not null NVARCHAR(512)"`
	V95             string `xorm:"not null NVARCHAR(512)"`
	Servicerecordid int64  `xorm:"BIGINT(8)"`
}

type Loginvalidation struct {
	Id                 int64     `xorm:"pk BIGINT(8)"`
	Agentid            string    `xorm:"NVARCHAR(510)"`
	Islogin            int       `xorm:"BIT(1)"`
	Loginip            string    `xorm:"NVARCHAR(510)"`
	Lastmarkeddatetime time.Time `xorm:"DATETIME(8)"`
}

type Servicerecord struct {
	Servicerecordid int64     `xorm:"not null pk BIGINT(8)"`
	Servicequeueid  []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Date            time.Time `xorm:"not null DATETIME(8)"`
}

type MtTasklist struct {
	Taskid                  int       `xorm:"not null pk INT(4)"`
	Tasktype                int       `xorm:"not null INT(4)"`
	Taskname                string    `xorm:"not null NVARCHAR(100)"`
	Campaignid              []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Sendtarget              int       `xorm:"not null default 0 INT(4)"`
	Issendbyemail           int       `xorm:"not null BIT(1)"`
	Emailsubjecttemplateid  int       `xorm:"INT(4)"`
	Emailtexttemplateid     int       `xorm:"INT(4)"`
	Issendbysms             int       `xorm:"not null BIT(1)"`
	Smstexttemplateid       int       `xorm:"INT(4)"`
	Issendbywebann          int       `xorm:"not null BIT(1)"`
	Webannsubjecttemplateid int       `xorm:"INT(4)"`
	Webanntexttemplateid    int       `xorm:"INT(4)"`
	Customparameter         string    `xorm:"NVARCHAR(8000)"`
	Lastchecktime           time.Time `xorm:"not null default 'getdate' DATETIME(8)"`
}

type Servicequeuetraffic struct {
	Servicequeuetrafficid int64     `xorm:"not null pk BIGINT(8)"`
	Trafficdate           time.Time `xorm:"DATETIME(8)"`
	Servicequeueid        []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	V0                    float32   `xorm:"FLOAT(8)"`
	V1                    float32   `xorm:"FLOAT(8)"`
	V2                    float32   `xorm:"FLOAT(8)"`
	V3                    float32   `xorm:"FLOAT(8)"`
	V4                    float32   `xorm:"FLOAT(8)"`
	V5                    float32   `xorm:"FLOAT(8)"`
	V6                    float32   `xorm:"FLOAT(8)"`
	V7                    float32   `xorm:"FLOAT(8)"`
	V8                    float32   `xorm:"FLOAT(8)"`
	V9                    float32   `xorm:"FLOAT(8)"`
	V10                   float32   `xorm:"FLOAT(8)"`
	V11                   float32   `xorm:"FLOAT(8)"`
	V12                   float32   `xorm:"FLOAT(8)"`
	V13                   float32   `xorm:"FLOAT(8)"`
	V14                   float32   `xorm:"FLOAT(8)"`
	V15                   float32   `xorm:"FLOAT(8)"`
	V16                   float32   `xorm:"FLOAT(8)"`
	V17                   float32   `xorm:"FLOAT(8)"`
	V18                   float32   `xorm:"FLOAT(8)"`
	V19                   float32   `xorm:"FLOAT(8)"`
	V20                   float32   `xorm:"FLOAT(8)"`
	V21                   float32   `xorm:"FLOAT(8)"`
	V22                   float32   `xorm:"FLOAT(8)"`
	V23                   float32   `xorm:"FLOAT(8)"`
	V24                   float32   `xorm:"FLOAT(8)"`
	V25                   float32   `xorm:"FLOAT(8)"`
	V26                   float32   `xorm:"FLOAT(8)"`
	V27                   float32   `xorm:"FLOAT(8)"`
	V28                   float32   `xorm:"FLOAT(8)"`
	V29                   float32   `xorm:"FLOAT(8)"`
	V30                   float32   `xorm:"FLOAT(8)"`
	V31                   float32   `xorm:"FLOAT(8)"`
	V32                   float32   `xorm:"FLOAT(8)"`
	V33                   float32   `xorm:"FLOAT(8)"`
	V34                   float32   `xorm:"FLOAT(8)"`
	V35                   float32   `xorm:"FLOAT(8)"`
	V36                   float32   `xorm:"FLOAT(8)"`
	V37                   float32   `xorm:"FLOAT(8)"`
	V38                   float32   `xorm:"FLOAT(8)"`
	V39                   float32   `xorm:"FLOAT(8)"`
	V40                   float32   `xorm:"FLOAT(8)"`
	V41                   float32   `xorm:"FLOAT(8)"`
	V42                   float32   `xorm:"FLOAT(8)"`
	V43                   float32   `xorm:"FLOAT(8)"`
	V44                   float32   `xorm:"FLOAT(8)"`
	V45                   float32   `xorm:"FLOAT(8)"`
	V46                   float32   `xorm:"FLOAT(8)"`
	V47                   float32   `xorm:"FLOAT(8)"`
	V48                   float32   `xorm:"FLOAT(8)"`
	V49                   float32   `xorm:"FLOAT(8)"`
	V50                   float32   `xorm:"FLOAT(8)"`
	V51                   float32   `xorm:"FLOAT(8)"`
	V52                   float32   `xorm:"FLOAT(8)"`
	V53                   float32   `xorm:"FLOAT(8)"`
	V54                   float32   `xorm:"FLOAT(8)"`
	V55                   float32   `xorm:"FLOAT(8)"`
	V56                   float32   `xorm:"FLOAT(8)"`
	V57                   float32   `xorm:"FLOAT(8)"`
	V58                   float32   `xorm:"FLOAT(8)"`
	V59                   float32   `xorm:"FLOAT(8)"`
	V60                   float32   `xorm:"FLOAT(8)"`
	V61                   float32   `xorm:"FLOAT(8)"`
	V62                   float32   `xorm:"FLOAT(8)"`
	V63                   float32   `xorm:"FLOAT(8)"`
	V64                   float32   `xorm:"FLOAT(8)"`
	V65                   float32   `xorm:"FLOAT(8)"`
	V66                   float32   `xorm:"FLOAT(8)"`
	V67                   float32   `xorm:"FLOAT(8)"`
	V68                   float32   `xorm:"FLOAT(8)"`
	V69                   float32   `xorm:"FLOAT(8)"`
	V70                   float32   `xorm:"FLOAT(8)"`
	V71                   float32   `xorm:"FLOAT(8)"`
	V72                   float32   `xorm:"FLOAT(8)"`
	V73                   float32   `xorm:"FLOAT(8)"`
	V74                   float32   `xorm:"FLOAT(8)"`
	V75                   float32   `xorm:"FLOAT(8)"`
	V76                   float32   `xorm:"FLOAT(8)"`
	V77                   float32   `xorm:"FLOAT(8)"`
	V78                   float32   `xorm:"FLOAT(8)"`
	V79                   float32   `xorm:"FLOAT(8)"`
	V80                   float32   `xorm:"FLOAT(8)"`
	V81                   float32   `xorm:"FLOAT(8)"`
	V82                   float32   `xorm:"FLOAT(8)"`
	V83                   float32   `xorm:"FLOAT(8)"`
	V84                   float32   `xorm:"FLOAT(8)"`
	V85                   float32   `xorm:"FLOAT(8)"`
	V86                   float32   `xorm:"FLOAT(8)"`
	V87                   float32   `xorm:"FLOAT(8)"`
	V88                   float32   `xorm:"FLOAT(8)"`
	V89                   float32   `xorm:"FLOAT(8)"`
	V90                   float32   `xorm:"FLOAT(8)"`
	V91                   float32   `xorm:"FLOAT(8)"`
	V92                   float32   `xorm:"FLOAT(8)"`
	V93                   float32   `xorm:"FLOAT(8)"`
	V94                   float32   `xorm:"FLOAT(8)"`
	V95                   float32   `xorm:"FLOAT(8)"`
	H0                    float32   `xorm:"FLOAT(8)"`
	H1                    float32   `xorm:"FLOAT(8)"`
	H2                    float32   `xorm:"FLOAT(8)"`
	H3                    float32   `xorm:"FLOAT(8)"`
	H4                    float32   `xorm:"FLOAT(8)"`
	H5                    float32   `xorm:"FLOAT(8)"`
	H6                    float32   `xorm:"FLOAT(8)"`
	H7                    float32   `xorm:"FLOAT(8)"`
	H8                    float32   `xorm:"FLOAT(8)"`
	H9                    float32   `xorm:"FLOAT(8)"`
	H10                   float32   `xorm:"FLOAT(8)"`
	H11                   float32   `xorm:"FLOAT(8)"`
	H12                   float32   `xorm:"FLOAT(8)"`
	H13                   float32   `xorm:"FLOAT(8)"`
	H14                   float32   `xorm:"FLOAT(8)"`
	H15                   float32   `xorm:"FLOAT(8)"`
	H16                   float32   `xorm:"FLOAT(8)"`
	H17                   float32   `xorm:"FLOAT(8)"`
	H18                   float32   `xorm:"FLOAT(8)"`
	H19                   float32   `xorm:"FLOAT(8)"`
	H20                   float32   `xorm:"FLOAT(8)"`
	H21                   float32   `xorm:"FLOAT(8)"`
	H22                   float32   `xorm:"FLOAT(8)"`
	H23                   float32   `xorm:"FLOAT(8)"`
	H24                   float32   `xorm:"FLOAT(8)"`
	H25                   float32   `xorm:"FLOAT(8)"`
	H26                   float32   `xorm:"FLOAT(8)"`
	H27                   float32   `xorm:"FLOAT(8)"`
	H28                   float32   `xorm:"FLOAT(8)"`
	H29                   float32   `xorm:"FLOAT(8)"`
	H30                   float32   `xorm:"FLOAT(8)"`
	H31                   float32   `xorm:"FLOAT(8)"`
	H32                   float32   `xorm:"FLOAT(8)"`
	H33                   float32   `xorm:"FLOAT(8)"`
	H34                   float32   `xorm:"FLOAT(8)"`
	H35                   float32   `xorm:"FLOAT(8)"`
	H36                   float32   `xorm:"FLOAT(8)"`
	H37                   float32   `xorm:"FLOAT(8)"`
	H38                   float32   `xorm:"FLOAT(8)"`
	H39                   float32   `xorm:"FLOAT(8)"`
	H40                   float32   `xorm:"FLOAT(8)"`
	H41                   float32   `xorm:"FLOAT(8)"`
	H42                   float32   `xorm:"FLOAT(8)"`
	H43                   float32   `xorm:"FLOAT(8)"`
	H44                   float32   `xorm:"FLOAT(8)"`
	H45                   float32   `xorm:"FLOAT(8)"`
	H46                   float32   `xorm:"FLOAT(8)"`
	H47                   float32   `xorm:"FLOAT(8)"`
	H48                   float32   `xorm:"FLOAT(8)"`
	H49                   float32   `xorm:"FLOAT(8)"`
	H50                   float32   `xorm:"FLOAT(8)"`
	H51                   float32   `xorm:"FLOAT(8)"`
	H52                   float32   `xorm:"FLOAT(8)"`
	H53                   float32   `xorm:"FLOAT(8)"`
	H54                   float32   `xorm:"FLOAT(8)"`
	H55                   float32   `xorm:"FLOAT(8)"`
	H56                   float32   `xorm:"FLOAT(8)"`
	H57                   float32   `xorm:"FLOAT(8)"`
	H58                   float32   `xorm:"FLOAT(8)"`
	H59                   float32   `xorm:"FLOAT(8)"`
	H60                   float32   `xorm:"FLOAT(8)"`
	H61                   float32   `xorm:"FLOAT(8)"`
	H62                   float32   `xorm:"FLOAT(8)"`
	H63                   float32   `xorm:"FLOAT(8)"`
	H64                   float32   `xorm:"FLOAT(8)"`
	H65                   float32   `xorm:"FLOAT(8)"`
	H66                   float32   `xorm:"FLOAT(8)"`
	H67                   float32   `xorm:"FLOAT(8)"`
	H68                   float32   `xorm:"FLOAT(8)"`
	H69                   float32   `xorm:"FLOAT(8)"`
	H70                   float32   `xorm:"FLOAT(8)"`
	H71                   float32   `xorm:"FLOAT(8)"`
	H72                   float32   `xorm:"FLOAT(8)"`
	H73                   float32   `xorm:"FLOAT(8)"`
	H74                   float32   `xorm:"FLOAT(8)"`
	H75                   float32   `xorm:"FLOAT(8)"`
	H76                   float32   `xorm:"FLOAT(8)"`
	H77                   float32   `xorm:"FLOAT(8)"`
	H78                   float32   `xorm:"FLOAT(8)"`
	H79                   float32   `xorm:"FLOAT(8)"`
	H80                   float32   `xorm:"FLOAT(8)"`
	H81                   float32   `xorm:"FLOAT(8)"`
	H82                   float32   `xorm:"FLOAT(8)"`
	H83                   float32   `xorm:"FLOAT(8)"`
	H84                   float32   `xorm:"FLOAT(8)"`
	H85                   float32   `xorm:"FLOAT(8)"`
	H86                   float32   `xorm:"FLOAT(8)"`
	H87                   float32   `xorm:"FLOAT(8)"`
	H88                   float32   `xorm:"FLOAT(8)"`
	H89                   float32   `xorm:"FLOAT(8)"`
	H90                   float32   `xorm:"FLOAT(8)"`
	H91                   float32   `xorm:"FLOAT(8)"`
	H92                   float32   `xorm:"FLOAT(8)"`
	H93                   float32   `xorm:"FLOAT(8)"`
	H94                   float32   `xorm:"FLOAT(8)"`
	H95                   float32   `xorm:"FLOAT(8)"`
	S0                    float32   `xorm:"FLOAT(8)"`
	S1                    float32   `xorm:"FLOAT(8)"`
	S2                    float32   `xorm:"FLOAT(8)"`
	S3                    float32   `xorm:"FLOAT(8)"`
	S4                    float32   `xorm:"FLOAT(8)"`
	S5                    float32   `xorm:"FLOAT(8)"`
	S6                    float32   `xorm:"FLOAT(8)"`
	S7                    float32   `xorm:"FLOAT(8)"`
	S8                    float32   `xorm:"FLOAT(8)"`
	S9                    float32   `xorm:"FLOAT(8)"`
	S10                   float32   `xorm:"FLOAT(8)"`
	S11                   float32   `xorm:"FLOAT(8)"`
	S12                   float32   `xorm:"FLOAT(8)"`
	S13                   float32   `xorm:"FLOAT(8)"`
	S14                   float32   `xorm:"FLOAT(8)"`
	S15                   float32   `xorm:"FLOAT(8)"`
	S16                   float32   `xorm:"FLOAT(8)"`
	S17                   float32   `xorm:"FLOAT(8)"`
	S18                   float32   `xorm:"FLOAT(8)"`
	S19                   float32   `xorm:"FLOAT(8)"`
	S20                   float32   `xorm:"FLOAT(8)"`
	S21                   float32   `xorm:"FLOAT(8)"`
	S22                   float32   `xorm:"FLOAT(8)"`
	S23                   float32   `xorm:"FLOAT(8)"`
	S24                   float32   `xorm:"FLOAT(8)"`
	S25                   float32   `xorm:"FLOAT(8)"`
	S26                   float32   `xorm:"FLOAT(8)"`
	S27                   float32   `xorm:"FLOAT(8)"`
	S28                   float32   `xorm:"FLOAT(8)"`
	S29                   float32   `xorm:"FLOAT(8)"`
	S30                   float32   `xorm:"FLOAT(8)"`
	S31                   float32   `xorm:"FLOAT(8)"`
	S32                   float32   `xorm:"FLOAT(8)"`
	S33                   float32   `xorm:"FLOAT(8)"`
	S34                   float32   `xorm:"FLOAT(8)"`
	S35                   float32   `xorm:"FLOAT(8)"`
	S36                   float32   `xorm:"FLOAT(8)"`
	S37                   float32   `xorm:"FLOAT(8)"`
	S38                   float32   `xorm:"FLOAT(8)"`
	S39                   float32   `xorm:"FLOAT(8)"`
	S40                   float32   `xorm:"FLOAT(8)"`
	S41                   float32   `xorm:"FLOAT(8)"`
	S42                   float32   `xorm:"FLOAT(8)"`
	S43                   float32   `xorm:"FLOAT(8)"`
	S44                   float32   `xorm:"FLOAT(8)"`
	S45                   float32   `xorm:"FLOAT(8)"`
	S46                   float32   `xorm:"FLOAT(8)"`
	S47                   float32   `xorm:"FLOAT(8)"`
	S48                   float32   `xorm:"FLOAT(8)"`
	S49                   float32   `xorm:"FLOAT(8)"`
	S50                   float32   `xorm:"FLOAT(8)"`
	S51                   float32   `xorm:"FLOAT(8)"`
	S52                   float32   `xorm:"FLOAT(8)"`
	S53                   float32   `xorm:"FLOAT(8)"`
	S54                   float32   `xorm:"FLOAT(8)"`
	S55                   float32   `xorm:"FLOAT(8)"`
	S56                   float32   `xorm:"FLOAT(8)"`
	S57                   float32   `xorm:"FLOAT(8)"`
	S58                   float32   `xorm:"FLOAT(8)"`
	S59                   float32   `xorm:"FLOAT(8)"`
	S60                   float32   `xorm:"FLOAT(8)"`
	S61                   float32   `xorm:"FLOAT(8)"`
	S62                   float32   `xorm:"FLOAT(8)"`
	S63                   float32   `xorm:"FLOAT(8)"`
	S64                   float32   `xorm:"FLOAT(8)"`
	S65                   float32   `xorm:"FLOAT(8)"`
	S66                   float32   `xorm:"FLOAT(8)"`
	S67                   float32   `xorm:"FLOAT(8)"`
	S68                   float32   `xorm:"FLOAT(8)"`
	S69                   float32   `xorm:"FLOAT(8)"`
	S70                   float32   `xorm:"FLOAT(8)"`
	S71                   float32   `xorm:"FLOAT(8)"`
	S72                   float32   `xorm:"FLOAT(8)"`
	S73                   float32   `xorm:"FLOAT(8)"`
	S74                   float32   `xorm:"FLOAT(8)"`
	S75                   float32   `xorm:"FLOAT(8)"`
	S76                   float32   `xorm:"FLOAT(8)"`
	S77                   float32   `xorm:"FLOAT(8)"`
	S78                   float32   `xorm:"FLOAT(8)"`
	S79                   float32   `xorm:"FLOAT(8)"`
	S80                   float32   `xorm:"FLOAT(8)"`
	S81                   float32   `xorm:"FLOAT(8)"`
	S82                   float32   `xorm:"FLOAT(8)"`
	S83                   float32   `xorm:"FLOAT(8)"`
	S84                   float32   `xorm:"FLOAT(8)"`
	S85                   float32   `xorm:"FLOAT(8)"`
	S86                   float32   `xorm:"FLOAT(8)"`
	S87                   float32   `xorm:"FLOAT(8)"`
	S88                   float32   `xorm:"FLOAT(8)"`
	S89                   float32   `xorm:"FLOAT(8)"`
	S90                   float32   `xorm:"FLOAT(8)"`
	S91                   float32   `xorm:"FLOAT(8)"`
	S92                   float32   `xorm:"FLOAT(8)"`
	S93                   float32   `xorm:"FLOAT(8)"`
	S94                   float32   `xorm:"FLOAT(8)"`
	S95                   float32   `xorm:"FLOAT(8)"`
	A0                    float32   `xorm:"FLOAT(8)"`
	A1                    float32   `xorm:"FLOAT(8)"`
	A2                    float32   `xorm:"FLOAT(8)"`
	A3                    float32   `xorm:"FLOAT(8)"`
	A4                    float32   `xorm:"FLOAT(8)"`
	A5                    float32   `xorm:"FLOAT(8)"`
	A6                    float32   `xorm:"FLOAT(8)"`
	A7                    float32   `xorm:"FLOAT(8)"`
	A8                    float32   `xorm:"FLOAT(8)"`
	A9                    float32   `xorm:"FLOAT(8)"`
	A10                   float32   `xorm:"FLOAT(8)"`
	A11                   float32   `xorm:"FLOAT(8)"`
	A12                   float32   `xorm:"FLOAT(8)"`
	A13                   float32   `xorm:"FLOAT(8)"`
	A14                   float32   `xorm:"FLOAT(8)"`
	A15                   float32   `xorm:"FLOAT(8)"`
	A16                   float32   `xorm:"FLOAT(8)"`
	A17                   float32   `xorm:"FLOAT(8)"`
	A18                   float32   `xorm:"FLOAT(8)"`
	A19                   float32   `xorm:"FLOAT(8)"`
	A20                   float32   `xorm:"FLOAT(8)"`
	A21                   float32   `xorm:"FLOAT(8)"`
	A22                   float32   `xorm:"FLOAT(8)"`
	A23                   float32   `xorm:"FLOAT(8)"`
	A24                   float32   `xorm:"FLOAT(8)"`
	A25                   float32   `xorm:"FLOAT(8)"`
	A26                   float32   `xorm:"FLOAT(8)"`
	A27                   float32   `xorm:"FLOAT(8)"`
	A28                   float32   `xorm:"FLOAT(8)"`
	A29                   float32   `xorm:"FLOAT(8)"`
	A30                   float32   `xorm:"FLOAT(8)"`
	A31                   float32   `xorm:"FLOAT(8)"`
	A32                   float32   `xorm:"FLOAT(8)"`
	A33                   float32   `xorm:"FLOAT(8)"`
	A34                   float32   `xorm:"FLOAT(8)"`
	A35                   float32   `xorm:"FLOAT(8)"`
	A36                   float32   `xorm:"FLOAT(8)"`
	A37                   float32   `xorm:"FLOAT(8)"`
	A38                   float32   `xorm:"FLOAT(8)"`
	A39                   float32   `xorm:"FLOAT(8)"`
	A40                   float32   `xorm:"FLOAT(8)"`
	A41                   float32   `xorm:"FLOAT(8)"`
	A42                   float32   `xorm:"FLOAT(8)"`
	A43                   float32   `xorm:"FLOAT(8)"`
	A44                   float32   `xorm:"FLOAT(8)"`
	A45                   float32   `xorm:"FLOAT(8)"`
	A46                   float32   `xorm:"FLOAT(8)"`
	A47                   float32   `xorm:"FLOAT(8)"`
	A48                   float32   `xorm:"FLOAT(8)"`
	A49                   float32   `xorm:"FLOAT(8)"`
	A50                   float32   `xorm:"FLOAT(8)"`
	A51                   float32   `xorm:"FLOAT(8)"`
	A52                   float32   `xorm:"FLOAT(8)"`
	A53                   float32   `xorm:"FLOAT(8)"`
	A54                   float32   `xorm:"FLOAT(8)"`
	A55                   float32   `xorm:"FLOAT(8)"`
	A56                   float32   `xorm:"FLOAT(8)"`
	A57                   float32   `xorm:"FLOAT(8)"`
	A58                   float32   `xorm:"FLOAT(8)"`
	A59                   float32   `xorm:"FLOAT(8)"`
	A60                   float32   `xorm:"FLOAT(8)"`
	A61                   float32   `xorm:"FLOAT(8)"`
	A62                   float32   `xorm:"FLOAT(8)"`
	A63                   float32   `xorm:"FLOAT(8)"`
	A64                   float32   `xorm:"FLOAT(8)"`
	A65                   float32   `xorm:"FLOAT(8)"`
	A66                   float32   `xorm:"FLOAT(8)"`
	A67                   float32   `xorm:"FLOAT(8)"`
	A68                   float32   `xorm:"FLOAT(8)"`
	A69                   float32   `xorm:"FLOAT(8)"`
	A70                   float32   `xorm:"FLOAT(8)"`
	A71                   float32   `xorm:"FLOAT(8)"`
	A72                   float32   `xorm:"FLOAT(8)"`
	A73                   float32   `xorm:"FLOAT(8)"`
	A74                   float32   `xorm:"FLOAT(8)"`
	A75                   float32   `xorm:"FLOAT(8)"`
	A76                   float32   `xorm:"FLOAT(8)"`
	A77                   float32   `xorm:"FLOAT(8)"`
	A78                   float32   `xorm:"FLOAT(8)"`
	A79                   float32   `xorm:"FLOAT(8)"`
	A80                   float32   `xorm:"FLOAT(8)"`
	A81                   float32   `xorm:"FLOAT(8)"`
	A82                   float32   `xorm:"FLOAT(8)"`
	A83                   float32   `xorm:"FLOAT(8)"`
	A84                   float32   `xorm:"FLOAT(8)"`
	A85                   float32   `xorm:"FLOAT(8)"`
	A86                   float32   `xorm:"FLOAT(8)"`
	A87                   float32   `xorm:"FLOAT(8)"`
	A88                   float32   `xorm:"FLOAT(8)"`
	A89                   float32   `xorm:"FLOAT(8)"`
	A90                   float32   `xorm:"FLOAT(8)"`
	A91                   float32   `xorm:"FLOAT(8)"`
	A92                   float32   `xorm:"FLOAT(8)"`
	A93                   float32   `xorm:"FLOAT(8)"`
	A94                   float32   `xorm:"FLOAT(8)"`
	A95                   float32   `xorm:"FLOAT(8)"`
	M0                    float32   `xorm:"FLOAT(8)"`
	M1                    float32   `xorm:"FLOAT(8)"`
	M2                    float32   `xorm:"FLOAT(8)"`
	M3                    float32   `xorm:"FLOAT(8)"`
	M4                    float32   `xorm:"FLOAT(8)"`
	M5                    float32   `xorm:"FLOAT(8)"`
	M6                    float32   `xorm:"FLOAT(8)"`
	M7                    float32   `xorm:"FLOAT(8)"`
	M8                    float32   `xorm:"FLOAT(8)"`
	M9                    float32   `xorm:"FLOAT(8)"`
	M10                   float32   `xorm:"FLOAT(8)"`
	M11                   float32   `xorm:"FLOAT(8)"`
	M12                   float32   `xorm:"FLOAT(8)"`
	M13                   float32   `xorm:"FLOAT(8)"`
	M14                   float32   `xorm:"FLOAT(8)"`
	M15                   float32   `xorm:"FLOAT(8)"`
	M16                   float32   `xorm:"FLOAT(8)"`
	M17                   float32   `xorm:"FLOAT(8)"`
	M18                   float32   `xorm:"FLOAT(8)"`
	M19                   float32   `xorm:"FLOAT(8)"`
	M20                   float32   `xorm:"FLOAT(8)"`
	M21                   float32   `xorm:"FLOAT(8)"`
	M22                   float32   `xorm:"FLOAT(8)"`
	M23                   float32   `xorm:"FLOAT(8)"`
	M24                   float32   `xorm:"FLOAT(8)"`
	M25                   float32   `xorm:"FLOAT(8)"`
	M26                   float32   `xorm:"FLOAT(8)"`
	M27                   float32   `xorm:"FLOAT(8)"`
	M28                   float32   `xorm:"FLOAT(8)"`
	M29                   float32   `xorm:"FLOAT(8)"`
	M30                   float32   `xorm:"FLOAT(8)"`
	M31                   float32   `xorm:"FLOAT(8)"`
	M32                   float32   `xorm:"FLOAT(8)"`
	M33                   float32   `xorm:"FLOAT(8)"`
	M34                   float32   `xorm:"FLOAT(8)"`
	M35                   float32   `xorm:"FLOAT(8)"`
	M36                   float32   `xorm:"FLOAT(8)"`
	M37                   float32   `xorm:"FLOAT(8)"`
	M38                   float32   `xorm:"FLOAT(8)"`
	M39                   float32   `xorm:"FLOAT(8)"`
	M40                   float32   `xorm:"FLOAT(8)"`
	M41                   float32   `xorm:"FLOAT(8)"`
	M42                   float32   `xorm:"FLOAT(8)"`
	M43                   float32   `xorm:"FLOAT(8)"`
	M44                   float32   `xorm:"FLOAT(8)"`
	M45                   float32   `xorm:"FLOAT(8)"`
	M46                   float32   `xorm:"FLOAT(8)"`
	M47                   float32   `xorm:"FLOAT(8)"`
	M48                   float32   `xorm:"FLOAT(8)"`
	M49                   float32   `xorm:"FLOAT(8)"`
	M50                   float32   `xorm:"FLOAT(8)"`
	M51                   float32   `xorm:"FLOAT(8)"`
	M52                   float32   `xorm:"FLOAT(8)"`
	M53                   float32   `xorm:"FLOAT(8)"`
	M54                   float32   `xorm:"FLOAT(8)"`
	M55                   float32   `xorm:"FLOAT(8)"`
	M56                   float32   `xorm:"FLOAT(8)"`
	M57                   float32   `xorm:"FLOAT(8)"`
	M58                   float32   `xorm:"FLOAT(8)"`
	M59                   float32   `xorm:"FLOAT(8)"`
	M60                   float32   `xorm:"FLOAT(8)"`
	M61                   float32   `xorm:"FLOAT(8)"`
	M62                   float32   `xorm:"FLOAT(8)"`
	M63                   float32   `xorm:"FLOAT(8)"`
	M64                   float32   `xorm:"FLOAT(8)"`
	M65                   float32   `xorm:"FLOAT(8)"`
	M66                   float32   `xorm:"FLOAT(8)"`
	M67                   float32   `xorm:"FLOAT(8)"`
	M68                   float32   `xorm:"FLOAT(8)"`
	M69                   float32   `xorm:"FLOAT(8)"`
	M70                   float32   `xorm:"FLOAT(8)"`
	M71                   float32   `xorm:"FLOAT(8)"`
	M72                   float32   `xorm:"FLOAT(8)"`
	M73                   float32   `xorm:"FLOAT(8)"`
	M74                   float32   `xorm:"FLOAT(8)"`
	M75                   float32   `xorm:"FLOAT(8)"`
	M76                   float32   `xorm:"FLOAT(8)"`
	M77                   float32   `xorm:"FLOAT(8)"`
	M78                   float32   `xorm:"FLOAT(8)"`
	M79                   float32   `xorm:"FLOAT(8)"`
	M80                   float32   `xorm:"FLOAT(8)"`
	M81                   float32   `xorm:"FLOAT(8)"`
	M82                   float32   `xorm:"FLOAT(8)"`
	M83                   float32   `xorm:"FLOAT(8)"`
	M84                   float32   `xorm:"FLOAT(8)"`
	M85                   float32   `xorm:"FLOAT(8)"`
	M86                   float32   `xorm:"FLOAT(8)"`
	M87                   float32   `xorm:"FLOAT(8)"`
	M88                   float32   `xorm:"FLOAT(8)"`
	M89                   float32   `xorm:"FLOAT(8)"`
	M90                   float32   `xorm:"FLOAT(8)"`
	M91                   float32   `xorm:"FLOAT(8)"`
	M92                   float32   `xorm:"FLOAT(8)"`
	M93                   float32   `xorm:"FLOAT(8)"`
	M94                   float32   `xorm:"FLOAT(8)"`
	M95                   float32   `xorm:"FLOAT(8)"`
}

type Licensekey struct {
	License string `xorm:"NVARCHAR(-1)"`
}

type Organizationgroup struct {
	Substitutionid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Organizationid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Timeoffdetail struct {
	Timeoffid   int       `xorm:"not null pk INT(4)"`
	Employeeid  []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Timeoffdate time.Time `xorm:"not null DATETIME(8)"`
	Isoutput    int       `xorm:"BIT(1)"`
}

type MtTemplate struct {
	Id           int    `xorm:"not null pk INT(4)"`
	Text         string `xorm:"not null NVARCHAR(8000)"`
	Templatetype int    `xorm:"not null INT(4)"`
	Templatename string `xorm:"not null NVARCHAR(100)"`
}

type Shrinkage struct {
	Id                 int     `xorm:"not null pk INT(4)"`
	Day                int     `xorm:"INT(4)"`
	H0                 float32 `xorm:"FLOAT(8)"`
	H1                 float32 `xorm:"FLOAT(8)"`
	H2                 float32 `xorm:"FLOAT(8)"`
	H3                 float32 `xorm:"FLOAT(8)"`
	H4                 float32 `xorm:"FLOAT(8)"`
	H5                 float32 `xorm:"FLOAT(8)"`
	H6                 float32 `xorm:"FLOAT(8)"`
	H7                 float32 `xorm:"FLOAT(8)"`
	H8                 float32 `xorm:"FLOAT(8)"`
	H9                 float32 `xorm:"FLOAT(8)"`
	H10                float32 `xorm:"FLOAT(8)"`
	H11                float32 `xorm:"FLOAT(8)"`
	H12                float32 `xorm:"FLOAT(8)"`
	H13                float32 `xorm:"FLOAT(8)"`
	H14                float32 `xorm:"FLOAT(8)"`
	H15                float32 `xorm:"FLOAT(8)"`
	H16                float32 `xorm:"FLOAT(8)"`
	H17                float32 `xorm:"FLOAT(8)"`
	H18                float32 `xorm:"FLOAT(8)"`
	H19                float32 `xorm:"FLOAT(8)"`
	H20                float32 `xorm:"FLOAT(8)"`
	H21                float32 `xorm:"FLOAT(8)"`
	H22                float32 `xorm:"FLOAT(8)"`
	H23                float32 `xorm:"FLOAT(8)"`
	Campaignscheduleid []byte  `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Acd struct {
	Acdid       []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Name        string `xorm:"NVARCHAR(510)"`
	Description string `xorm:"NVARCHAR(510)"`
}

type Swaperrlog struct {
	Id        int       `xorm:"not null INT(4)"`
	Createdat time.Time `xorm:"not null DATETIME(8)"`
	Createdid string    `xorm:"not null VARCHAR(20)"`
	Org1      string    `xorm:"NVARCHAR(40)"`
	Org2      string    `xorm:"NVARCHAR(40)"`
	Agentname string    `xorm:"not null NVARCHAR(100)"`
	Agentid   string    `xorm:"not null VARCHAR(10)"`
	Loginid   string    `xorm:"VARCHAR(10)"`
	Swapat    time.Time `xorm:"DATETIME(8)"`
	Message   string    `xorm:"NVARCHAR(200)"`
}

type Absence struct {
	Id                   int       `xorm:"not null pk INT(4)"`
	Employeeid           []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Campaignid           []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Startat              time.Time `xorm:"not null DATETIME(8)"`
	Length               int       `xorm:"not null INT(4)"`
	Status               string    `xorm:"not null CHAR(1)"`
	Shrinkagetypeid      int       `xorm:"INT(4)"`
	Assignmentid         int64     `xorm:"BIGINT(8)"`
	Targetdate           time.Time `xorm:"DATETIME(8)"`
	Modified             time.Time `xorm:"DATETIME(8)"`
	Modifiedby           []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Note                 string    `xorm:"NVARCHAR(500)"`
	Absencetypeid        int       `xorm:"INT(4)"`
	Shrinkagetime        int       `xorm:"INT(4)"`
	Createdat            time.Time `xorm:"not null default 'getdate' DATETIME(8)"`
	Iscompleted          int       `xorm:"not null default 1 BIT(1)"`
	Merageassignmentname string    `xorm:"NVARCHAR(400)"`
	Merageassignmenttime string    `xorm:"CHAR(39)"`
	Absenceparts         string    `xorm:"VARCHAR(1000)"`
	Termids              string    `xorm:"VARCHAR(100)"`
	Org1                 []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
	Org2                 []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
}

type Organizationmanager struct {
	Managerid      int    `xorm:"not null pk INT(4)"`
	Employeeid     []byte `xorm:"unique(UQ__tmp_ms_x__467DFFA17B113988) UNIQUEIDENTIFIER(16)"`
	Organizationid []byte `xorm:"unique(UQ__tmp_ms_x__467DFFA17B113988) UNIQUEIDENTIFIER(16)"`
}

type Absencetype struct {
	Id        int       `xorm:"not null pk INT(4)"`
	Name      string    `xorm:"not null NVARCHAR(100)"`
	Isdeleted int       `xorm:"not null default 0 BIT(1)"`
	Deleteat  time.Time `xorm:"DATETIME(8)"`
}

type Timebox struct {
	Timeboxid []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
}

type Absencehistory struct {
	Id            int       `xorm:"not null pk INT(4)"`
	Absenceid     int       `xorm:"not null INT(4)"`
	Absencetypeid int       `xorm:"not null INT(4)"`
	Note          string    `xorm:"NVARCHAR(100)"`
	Isshrinkage   int       `xorm:"not null BIT(1)"`
	Createdby     []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Createdat     time.Time `xorm:"not null default 'getdate' DATETIME(8)"`
}

type Absencemessagetransformcheckinfo struct {
	Campaignid    []byte `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Istomanager   int    `xorm:"not null BIT(1)"`
	Istoagent     int    `xorm:"not null BIT(1)"`
	Issendbyemail int    `xorm:"not null BIT(1)"`
	Issendsms     int    `xorm:"not null BIT(1)"`
	Sendstatus    string `xorm:"not null CHAR(1)"`
	Templateid    int    `xorm:"not null INT(4)"`
}

type Agentscore struct {
	Id          int       `xorm:"not null pk INT(4)"`
	Agentid     []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Scorerid    []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Scoretypeid int       `xorm:"not null INT(4)"`
	Score       float32   `xorm:"not null FLOAT(8)"`
	Scoredate   time.Time `xorm:"not null default 'getdate' DATETIME(8)"`
	Note        string    `xorm:"NVARCHAR(1000)"`
	Happenat    time.Time `xorm:"DATETIME(8)"`
}

type Termstyle struct {
	Termstyleid                    []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Childclasstype                 string `xorm:"not null NVARCHAR(510)"`
	Occupied                       int    `xorm:"not null BIT(1)"`
	Onservice                      int    `xorm:"not null BIT(1)"`
	Asarest                        int    `xorm:"BIT(1)"`
	Inuse                          int    `xorm:"BIT(1)"`
	Startvalue                     int    `xorm:"not null INT(4)"`
	Endvalue                       int    `xorm:"not null INT(4)"`
	Background                     string `xorm:"NVARCHAR(510)"`
	Foreground                     string `xorm:"NVARCHAR(510)"`
	Name                           string `xorm:"NVARCHAR(510)"`
	Type                           string `xorm:"NVARCHAR(510)"`
	Label                          string `xorm:"NVARCHAR(510)"`
	Asawork                        int    `xorm:"BIT(1)"`
	Gapguaranteed                  int    `xorm:"BIT(1)"`
	Ignoreadherence                int    `xorm:"BIT(1)"`
	Country                        string `xorm:"NVARCHAR(510)"`
	Timezoneid                     string `xorm:"NVARCHAR(510)"`
	Customtag                      string `xorm:"NVARCHAR(510)"`
	Servicequeueid                 []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Estimationpriority             int    `xorm:"INT(4)"`
	Template1id                    []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Template2id                    []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Template3id                    []byte `xorm:"UNIQUEIDENTIFIER(16)"`
	Validworkingweekdayfromsun2sat int    `xorm:"INT(4)"`
	Validdayoffweekdayfromsun2sat  int    `xorm:"INT(4)"`
	Validdayoffmonthday            int    `xorm:"INT(4)"`
	Shifttimetag                   string `xorm:"NVARCHAR(510)"`
}

type AgentbiddingTimeoff struct {
	Employeeid            []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Biddingsettingid      []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Biddingannouncementid []byte `xorm:"not null UNIQUEIDENTIFIER(16)"`
}

type Agentassignmentsharedquota struct {
	Agentassignmentsharedquotaid int64   `xorm:"not null pk BIGINT(8)"`
	Attendanceid                 int64   `xorm:"BIGINT(8)"`
	Assignmenttypeid             []byte  `xorm:"UNIQUEIDENTIFIER(16)"`
	Quota                        float32 `xorm:"FLOAT(8)"`
}

type AgentbiddingDayoff struct {
	Employeeid            []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Biddingsettingid      []byte `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Biddingannouncementid []byte `xorm:"not null UNIQUEIDENTIFIER(16)"`
}

type Swapsubevent struct {
	Swapsubeventid      []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Applicationdate     time.Time `xorm:"not null DATETIME(8)"`
	Exchangedate        time.Time `xorm:"not null DATETIME(8)"`
	Employeeaid         []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Employeeainfo       string    `xorm:"not null NVARCHAR(128)"`
	Assignmentaid       int64     `xorm:"not null BIGINT(8)"`
	Assignmentainfo     string    `xorm:"not null NVARCHAR(256)"`
	Subeventaid         int64     `xorm:"not null BIGINT(8)"`
	Subeventainfo       string    `xorm:"not null NVARCHAR(256)"`
	Descriptiona        string    `xorm:"NVARCHAR(510)"`
	Employeebid         []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Employeebinfo       string    `xorm:"not null NVARCHAR(128)"`
	Assignmentbid       int64     `xorm:"not null BIGINT(8)"`
	Assignmentbinfo     string    `xorm:"not null NVARCHAR(256)"`
	Subeventbid         int64     `xorm:"not null BIGINT(8)"`
	Subeventbinfo       string    `xorm:"not null NVARCHAR(256)"`
	Descriptionb        string    `xorm:"NVARCHAR(510)"`
	Status              string    `xorm:"not null NVARCHAR(32)"`
	Completedate        time.Time `xorm:"DATETIME(8)"`
	Maturitydate        time.Time `xorm:"not null DATETIME(8)"`
	Swaptype            string    `xorm:"not null NVARCHAR(32)"`
	Substitutiongroupid []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Warningmessagea     string    `xorm:"NVARCHAR(-1)"`
	Warningmessageb     string    `xorm:"NVARCHAR(-1)"`
	Version             time.Time `xorm:"TIMESTAMP(8)"`
}

type Offdutesetting struct {
	Offdutesettingid []byte    `xorm:"not null pk UNIQUEIDENTIFIER(16)"`
	Offdutegroupid   []byte    `xorm:"not null UNIQUEIDENTIFIER(16)"`
	Offdutedate      time.Time `xorm:"DATETIME(8)"`
	Offdutetype      string    `xorm:"VARCHAR(10)"`
	Offdutelimite    int       `xorm:"INT(4)"`
}

type Forecast struct {
	Forecastid         int       `xorm:"not null pk INT(4)"`
	Name               string    `xorm:"not null NVARCHAR(60)"`
	Description        string    `xorm:"NVARCHAR(200)"`
	Startat            time.Time `xorm:"not null DATETIME(8)"`
	Endat              time.Time `xorm:"not null DATETIME(8)"`
	Servicesecond      int       `xorm:"INT(4)"`
	Abandonrate        int       `xorm:"not null INT(4)"`
	Servicelevelgoalid int       `xorm:"INT(4)"`
	Servicequeueid     []byte    `xorm:"UNIQUEIDENTIFIER(16)"`
}
