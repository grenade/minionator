package main

type Bug struct {
  Id int `json:"id"`
  Alias string `json:"alias"`
  Url string `json:"url"`
  Summary string `json:"summary"`
  Platform string `json:"platform"`
  Creator string `json:"creator"`
  LastChangeTime string `json:"last_change_time"`
  CcList []string `json:"cc"`
  AssignedTo string `json:"assigned_to"`
  Whiteboard string `json:"whiteboard"`
  CreationTime string `json:"creation_time"`
  DependsOn []int `json:"depends_on"`
  Resolution string `json:"resolution"`
  OpSys string `json:"op_sys"`
  Status string `json:"status"`
  IsOpen bool `json:"is_open"`
  Severity string `json:"severity"`
  Component string `json:"component"`
  Product string `json:"product"`
}

type BugsApiResponse struct {
  Bugs []Bug `json:"bugs"`
  Faults []string `json:"faults"`
}

type Bugs []Bug

type DependsOnAppender struct {
  Add []int `json:"add"`
}

type Comment struct {
  Body string `json:"body"`
  IsPrivate bool `json:"is_private"`
  IsMarkdown bool `json:"is_markdown"`
}

type ReOpenChildMessage struct {
  Ids []int `json:"ids"`
  Status string `json:"status"`
  DependsOn DependsOnAppender `json:"depends_on"`
  Comment Comment `json:"comment"`
}

type CloseChildMessage struct {
  Ids []int `json:"ids"`
  Status string `json:"status"`
  Resolution string `json:"resolution"`
  Comment Comment `json:"comment"`
}