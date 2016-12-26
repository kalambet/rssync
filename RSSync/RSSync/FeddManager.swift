//
//  FeddManager.swift
//  RSSync
//
//  Created by Peter Kalambet on 11/12/2016.
//  Copyright © 2016 Peter Kalambet. All rights reserved.
//

import Foundation

class FeedManager {
    
    //MARK: Singleton implementation
    
    static let sharedInstance : FeedManager = {
        let instance = FeedManager()
        return instance
    }()
    
    private init() {}
    
    // MARK: Type properites
    
    private var feedsList: [Feed] = []
    
    var feeds: [Feed] {
        return feedsList
    }
    
    func loadFeeds(provider: FeedLoaderProvider, callback: (_ err: Error) -> Void) {
        provider.Load(callback: callback)
    }
}
