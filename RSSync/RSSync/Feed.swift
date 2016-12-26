//
//  Feed.swift
//  RSSync
//
//  Created by Peter Kalambet on 11/12/2016.
//  Copyright © 2016 Peter Kalambet. All rights reserved.
//

import Foundation

class Feed {
    var title: String = ""
    var description: String?
    var link: NSURL = NSURL()
    var copyraight: String?
    var atomLink: NSURL?
    var lastUpdated: NSDate = NSDate()
    var items: [FeedItem] = []
}
