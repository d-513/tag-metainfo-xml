//    Copyright (C) 2021 dada513
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <https://www.gnu.org/licenses/>.
package main

import "encoding/xml"

type Provides struct {
	Id string `xml:"id"`
}
type Url struct {
	Url  string `xml:",innerxml"`
	Type string `xml:"type,attr"`
}

type ContentRatingContentAttribute struct {
	Id   string `xml:"id,attr"`
	Type string `xml:",innerxml"`
}

type ContentRating struct {
	Type             string                          `xml:"type,attr"`
	ContentAttribute []ContentRatingContentAttribute `xml:"content_attribute"`
}

type Release struct {
	Version string `xml:"version,attr"`
	Date    string `xml:"date,attr"`
}
type Releases struct {
	Release []Release `xml:"release"`
}

type Screenshot struct {
	Type    string `xml:"type,attr"`
	Caption string `xml:"caption"`
	Image   ScreenshotImage `xml:"image"`
}

type ScreenshotImage struct {
	Type   string `xml:"type,attr"`
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Url    string `xml:",innerxml"`
}

type Tag struct {
	XMLName xml.Name
	Content string `xml:",innerxml"`
}

type Description struct {
	Items []Tag `xml:",any"`
}

type Launchable struct {
	Type string `xml:"type,attr"`
	DesktopId string `xml:",innerxml"`
}

type component struct {
	Type            string        `xml:"type,attr"`
	Id              string        `xml:"id"`
	Provides        Provides      `xml:"provides"`
	Launchable		Launchable	  `xml:"launchable"`
	Name            string        `xml:"name"`
	DevName         string        `xml:"developer_name"`
	Summary         string        `xml:"summary"`
	MetadataLicense string        `xml:"metadata_license"`
	ProjectLicense  string        `xml:"project_license"`
	Url             []Url         `xml:"url"`
	Description     Description   `xml:"description"`
	Screenshots     []Screenshot  `xml:"screenshots>screenshot"`
	Releases        Releases      `xml:"releases"`
	ContentRating   ContentRating `xml:"content_rating"`
}
