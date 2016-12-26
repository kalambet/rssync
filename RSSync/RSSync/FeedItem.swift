//
//  FeedItem.swift
//  RSSync
//
//  Created by Peter Kalambet on 11/12/2016.
//  Copyright © 2016 Peter Kalambet. All rights reserved.
//

import Foundation

class FeedItem {
    var titile: String = ""
    var description: String?
    var link: NSURL = NSURL()
    var guid: NSURL = NSURL()
    var pubDate: NSDate = NSDate()
}
