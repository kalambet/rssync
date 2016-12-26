//
//  FeedLoaderProvider.swift
//  RSSync
//
//  Created by Peter Kalambet on 11/12/2016.
//  Copyright © 2016 Peter Kalambet. All rights reserved.
//

import Foundation

protocol FeedLoaderProvider {
    func Load(callback: (_ err: Error) -> Void) -> Void
}
