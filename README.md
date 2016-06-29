[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-collector-osv.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-collector-osv )
[![Go Report Card](http://goreportcard.com/badge/intelsdi-x/snap-plugin-collector-osv)](http://goreportcard.com/report/intelsdi-x/snap-plugin-collector-osv)
# snap collector plugin - osv

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

### System Requirements

#### OSv host with rest api support. For dev builds please run OSv with external network support

```
sudo ./scripts/run.py -n -v --api
```


#### Compile plugin
```
make
```

### Documentation

### Examples
Example running osv, passthru processor, and writing data to a file.

In one terminal window, open the snap daemon :
```
$ snapd -l 1
```

In another terminal window:
Load osv plugin
```
$ snapctl plugin load $SNAP_OSV_PLUGIN_DIR/build/rootfs/snap-plugin-collector-osv
```
See available metrics for your system
```
$ snapctl metric list
```

Create a task JSON file:    
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/osv/trace/wait/waitqueue_wake_one": {},
                "/intel/osv/trace/callout/callout_reset": {},
                "/intel/osv/cpu/cputime": {},
                "/intel/osv/memory/free": {}
            },
            "config": {
                "/intel/osv": {
                    "swag_ip": "192.168.122.89",
                    "swag_port": 8000
                }
            },
            "process": [
                {
                    "plugin_name": "passthru",
                    "process": null,
                    "publish": [
                        {                         
                            "plugin_name": "file",
                            "config": {
                                "file": "/tmp/published_psutil"
                            }
                        }
                    ],
                    "config": null
                }
            ],
            "publish": null
        }
    }
}
```

Load passthru plugin for processing:
```
$ snapctl plugin load build/rootfs/plugin/snap-processor-passthru
Plugin loaded
Name: passthru
Version: 1
Type: processor
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:44:03 PST
```

Load file plugin for publishing:
```
$ snapctl plugin load build/rootfs/plugin/snap-publisher-file
Plugin loaded
Name: file
Version: 3
Type: publisher
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:41:39 PST
```

Change ip address and port of osv host in task manifest:
```
vim $SNAP_OSV_PLUGIN_DIR/example/osv-file-example.json
```

Create task:
```
$ snapctl task create -t $SNAP_OSV_PLUGIN_DIR/example/osv-file-example.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See file output (this is just part of the file):

|NAMESPACE|DATA|TIMESTAMP|SOURCE
|---|---|---|---|
|/intel/osv/cpu/cputime|176521305|2015-11-25 15:36:04.225846442 +0000 UTC|192.168.122.89|
|/intel/osv/memory/free|2023403520|2015-11-25 15:36:04.226192641 +0000 UTC|192.168.122.89|
|/intel/osv/trace/callout/callout_reset|206217|2015-11-25 15:36:04.226534352 +0000 UTC|192.168.122.89|
|/intel/osv/trace/wait/waitqueue_wake_one|1319942|2015-11-25 15:36:04.226810341 +0000 UTC|192.168.122.89|


### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Data Type | Description
----------|-----------|-----------------------
/intel/osv/cpu/cputime| uint64| cputime
/intel/osv/memory/free| uint64| free memory
/intel/osv/memory/total| uint64| total memory
/intel/osv/trace/virtio/virtio_wait_for_queue| uint64|
/intel/osv/trace/virtio/virtio_enable_interrupts| uint64|
/intel/osv/trace/virtio/virtio_disable_interrupts| uint64|
/intel/osv/trace/virtio/virtio_kicked_event_idx| uint64|
/intel/osv/trace/virtio/virtio_add_buf| uint64|
/intel/osv/trace/virtio/virtio_net_rx_packet| uint64|
/intel/osv/trace/virtio/virtio_net_rx_wake| uint64|
/intel/osv/trace/virtio/virtio_net_fill_rx_ring| uint64|
/intel/osv/trace/virtio/virtio_net_fill_rx_ring_added| uint64|
/intel/osv/trace/virtio/virtio_net_tx_packet| uint64|
/intel/osv/trace/virtio/virtio_net_tx_failed_add_buf| uint64|
/intel/osv/trace/virtio/virtio_net_tx_no_space_calling_gc| uint64|
/intel/osv/trace/virtio/virtio_net_tx_packet_size| uint64|
/intel/osv/trace/virtio/virtio_net_tx_xmit_one_failed_to_post| uint64|
/intel/osv/trace/virtio/virtio_blk_read_config_capacity| uint64|
/intel/osv/trace/virtio/virtio_blk_read_config_size_max| uint64|
/intel/osv/trace/virtio/virtio_blk_read_config_seg_max| uint64|
/intel/osv/trace/virtio/virtio_blk_read_config_geometry| uint64|
/intel/osv/trace/virtio/virtio_blk_read_config_blk_siz| uint64|
/intel/osv/trace/virtio/virtio_blk_read_config_topology| uint64|
/intel/osv/trace/virtio/virtio_blk_read_config_wce| uint64|
/intel/osv/trace/virtio/virtio_blk_read_config_ro| uint64|
/intel/osv/trace/virtio/virtio_blk_make_request_seg_max| uint64|
/intel/osv/trace/virtio/virtio_blk_make_request_readonly| uint64|
/intel/osv/trace/virtio/virtio_blk_wake| uint64|
/intel/osv/trace/virtio/virtio_blk_strategy| uint64|
/intel/osv/trace/virtio/virtio_blk_req_ok| uint64|
/intel/osv/trace/virtio/virtio_blk_req_unsupp| uint64|
/intel/osv/trace/virtio/virtio_blk_req_err| uint64|
/intel/osv/trace/net/net_packet_in| uint64|
/intel/osv/trace/net/net_packet_out| uint64|
/intel/osv/trace/net/net_packet_handling| uint64|
/intel/osv/trace/tcp/tcp_state| uint64|
/intel/osv/trace/tcp/tcp_input_ack| uint64|
/intel/osv/trace/tcp/tcp_output| uint64|
/intel/osv/trace/tcp/tcp_output_error| uint64|
/intel/osv/trace/tcp/tcp_output_resched_start| uint64|
/intel/osv/trace/tcp/tcp_output_resched_end| uint64|
/intel/osv/trace/tcp/tcp_output_start| uint64|
/intel/osv/trace/tcp/tcp_output_ret| uint64|
/intel/osv/trace/tcp/tcp_output_just_ret| uint64|
/intel/osv/trace/tcp/tcp_output_cant_take_inp_lock| uint64|
/intel/osv/trace/tcp/tcp_timer_tso_flush| uint64|
/intel/osv/trace/tcp/tcp_timer_tso_flush_ret| uint64|
/intel/osv/trace/tcp/tcp_timer_tso_flush_err| uint64|
/intel/osv/trace/memory/memory_malloc| uint64|
/intel/osv/trace/memory/memory_malloc_mempool| uint64|
/intel/osv/trace/memory/memory_malloc_large| uint64|
/intel/osv/trace/memory/memory_malloc_page| uint64|
/intel/osv/trace/memory/memory_free| uint64|
/intel/osv/trace/memory/memory_realloc| uint64|
/intel/osv/trace/memory/memory_page_alloc| uint64|
/intel/osv/trace/memory/memory_page_free| uint64|
/intel/osv/trace/memory/memory_huge_failure| uint64|
/intel/osv/trace/memory/memory_reclaim| uint64|
/intel/osv/trace/memory/memory_wait| uint64|
/intel/osv/trace/memory/memory_mmap| uint64|
/intel/osv/trace/memory/memory_mmap_err| uint64|
/intel/osv/trace/memory/memory_mmap_ret| uint64|
/intel/osv/trace/memory/memory_munmap| uint64|
/intel/osv/trace/memory/memory_munmap_err| uint64|
/intel/osv/trace/memory/memory_munmap_ret| uint64|
/intel/osv/trace/callout/callout_init| uint64|
/intel/osv/trace/callout/callout_reset| uint64|
/intel/osv/trace/callout/callout_stop_wait| uint64|
/intel/osv/trace/callout/callout_stop| uint64|
/intel/osv/trace/callout/callout_thread_waiting| uint64|
/intel/osv/trace/callout/callout_thread_dispatching| uint64|
/intel/osv/trace/wait/waitqueue_wait| uint64|
/intel/osv/trace/wait/waitqueue_wake_one| uint64|
/intel/osv/trace/wait/waitqueue_wake_all| uint64|
/intel/osv/trace/anync/async_timer_task_create| uint64|
/intel/osv/trace/anync/async_timer_task_destroy| uint64|
/intel/osv/trace/anync/async_timer_task_reschedule| uint64|
/intel/osv/trace/anync/async_timer_task_cancel| uint64|
/intel/osv/trace/anync/async_timer_task_shutdown| uint64|
/intel/osv/trace/anync/async_timer_task_fire| uint64|
/intel/osv/trace/anync/async_timer_task_misfire| uint64|
/intel/osv/trace/anync/async_timer_task_insert| uint64|
/intel/osv/trace/anync/async_timer_task_remove| uint64|
/intel/osv/trace/anync/async_worker_started| uint64|
/intel/osv/trace/anync/async_worker_timer_fire| uint64|
/intel/osv/trace/anync/async_worker_timer_fire_ret| uint64|
/intel/osv/trace/anync/async_worker_fire| uint64|
/intel/osv/trace/anync/async_worker_fire_ret| uint64|
/intel/osv/trace/vfs/vfs_open| uint64|
/intel/osv/trace/vfs/vfs_open_ret| uint64|
/intel/osv/trace/vfs/vfs_open_err| uint64|
/intel/osv/trace/vfs/vfs_close| uint64|
/intel/osv/trace/vfs/vfs_close_ret| uint64|
/intel/osv/trace/vfs/vfs_close_err| uint64|
/intel/osv/trace/vfs/vfs_mknod| uint64|
/intel/osv/trace/vfs/vfs_mknod_ret| uint64|
/intel/osv/trace/vfs/vfs_mknod_err| uint64|
/intel/osv/trace/vfs/vfs_lseek| uint64|
/intel/osv/trace/vfs/vfs_lseek_ret| uint64|
/intel/osv/trace/vfs/vfs_lseek_err| uint64|
/intel/osv/trace/vfs/vfs_pread| uint64|
/intel/osv/trace/vfs/vfs_pread_ret| uint64|
/intel/osv/trace/vfs/vfs_pread_err| uint64|
/intel/osv/trace/vfs/vfs_pwrite| uint64|
/intel/osv/trace/vfs/vfs_pwrite_ret| uint64|
/intel/osv/trace/vfs/vfs_pwrite_err| uint64|
/intel/osv/trace/vfs/vfs_pwritev| uint64|
/intel/osv/trace/vfs/vfs_pwritev_ret| uint64|
/intel/osv/trace/vfs/vfs_pwritev_err| uint64|
/intel/osv/trace/vfs/vfs_ioctl| uint64|
/intel/osv/trace/vfs/vfs_ioctl_ret| uint64|
/intel/osv/trace/vfs/vfs_ioctl_err| uint64|
/intel/osv/trace/vfs/vfs_fsync| uint64|
/intel/osv/trace/vfs/vfs_fsync_ret| uint64|
/intel/osv/trace/vfs/vfs_fsync_err| uint64|
/intel/osv/trace/vfs/vfs_fstat| uint64|
/intel/osv/trace/vfs/vfs_fstat_ret| uint64|
/intel/osv/trace/vfs/vfs_fstat_err| uint64|
/intel/osv/trace/vfs/vfs_readdir| uint64|
/intel/osv/trace/vfs/vfs_readdir_ret| uint64|
/intel/osv/trace/vfs/vfs_readdir_err| uint64|
/intel/osv/trace/vfs/vfs_mkdir| uint64|
/intel/osv/trace/vfs/vfs_mkdir_ret| uint64|
/intel/osv/trace/vfs/vfs_mkdir_err| uint64|
/intel/osv/trace/vfs/vfs_rmdir| uint64|
/intel/osv/trace/vfs/vfs_rmdir_ret| uint64|
/intel/osv/trace/vfs/vfs_rmdir_err| uint64|
/intel/osv/trace/vfs/vfs_rename| uint64|
/intel/osv/trace/vfs/vfs_rename_ret| uint64|
/intel/osv/trace/vfs/vfs_rename_err| uint64|
/intel/osv/trace/vfs/vfs_chdir| uint64|
/intel/osv/trace/vfs/vfs_chdir_ret| uint64|
/intel/osv/trace/vfs/vfs_fchdir| uint64|
/intel/osv/trace/vfs/vfs_fchdir_ret| uint64|
/intel/osv/trace/vfs/vfs_fchdir_err| uint64|
/intel/osv/trace/vfs/vfs_link| uint64|
/intel/osv/trace/vfs/vfs_link_ret| uint64|
/intel/osv/trace/vfs/vfs_link_err| uint64|
/intel/osv/trace/vfs/vfs_symlink| uint64|
/intel/osv/trace/vfs/vfs_symlink_ret| uint64|
/intel/osv/trace/vfs/vfs_symlink_err| uint64|
/intel/osv/trace/vfs/vfs_unlink| uint64|
/intel/osv/trace/vfs/vfs_unlink_ret| uint64|
/intel/osv/trace/vfs/vfs_unlink_err| uint64|
/intel/osv/trace/vfs/vfs_stat| uint64|
/intel/osv/trace/vfs/vfs_stat_ret| uint64|
/intel/osv/trace/vfs/vfs_stat_err| uint64|
/intel/osv/trace/vfs/vfs_lstat| uint64|
/intel/osv/trace/vfs/vfs_lstat_ret| uint64|
/intel/osv/trace/vfs/vfs_lstat_err| uint64|
/intel/osv/trace/vfs/vfs_statfs| uint64|
/intel/osv/trace/vfs/vfs_statfs_ret| uint64|
/intel/osv/trace/vfs/vfs_statfs_err| uint64|
/intel/osv/trace/vfs/vfs_fstatfs| uint64|
/intel/osv/trace/vfs/vfs_fstatfs_ret| uint64|
/intel/osv/trace/vfs/vfs_fstatfs_err| uint64|
/intel/osv/trace/vfs/vfs_getcwd| uint64|
/intel/osv/trace/vfs/vfs_getcwd_ret| uint64|
/intel/osv/trace/vfs/vfs_getcwd_err| uint64|
/intel/osv/trace/vfs/vfs_dup| uint64|
/intel/osv/trace/vfs/vfs_dup_ret| uint64|
/intel/osv/trace/vfs/vfs_dup_err| uint64|
/intel/osv/trace/vfs/vfs_dup3| uint64|
/intel/osv/trace/vfs/vfs_dup3_ret| uint64|
/intel/osv/trace/vfs/vfs_dup3_err| uint64|
/intel/osv/trace/vfs/vfs_fcntl| uint64|
/intel/osv/trace/vfs/vfs_fcntl_ret| uint64|
/intel/osv/trace/vfs/vfs_fcntl_err| uint64|
/intel/osv/trace/vfs/vfs_access| uint64|
/intel/osv/trace/vfs/vfs_access_ret| uint64|
/intel/osv/trace/vfs/vfs_access_err| uint64|
/intel/osv/trace/vfs/vfs_isatty| uint64|
/intel/osv/trace/vfs/vfs_isatty_ret| uint64|
/intel/osv/trace/vfs/vfs_isatty_err| uint64|
/intel/osv/trace/vfs/vfs_truncate| uint64|
/intel/osv/trace/vfs/vfs_truncate_ret| uint64|
/intel/osv/trace/vfs/vfs_truncate_err| uint64|
/intel/osv/trace/vfs/vfs_ftruncate| uint64|
/intel/osv/trace/vfs/vfs_ftruncate_ret| uint64|
/intel/osv/trace/vfs/vfs_ftruncate_err| uint64|
/intel/osv/trace/vfs/vfs_fallocate| uint64|
/intel/osv/trace/vfs/vfs_fallocate_ret| uint64|
/intel/osv/trace/vfs/vfs_fallocate_err| uint64|
/intel/osv/trace/vfs/vfs_utimes| uint64|
/intel/osv/trace/vfs/vfs_utimes_ret| uint64|
/intel/osv/trace/vfs/vfs_utimes_err| uint64|
/intel/osv/trace/vfs/vfs_utimensat| uint64|
/intel/osv/trace/vfs/vfs_utimensat_ret| uint64|
/intel/osv/trace/vfs/vfs_utimensat_err| uint64|
/intel/osv/trace/vfs/vfs_futimens| uint64|
/intel/osv/trace/vfs/vfs_futimens_ret| uint64|
/intel/osv/trace/vfs/vfs_futimens_err| uint64|
/intel/osv/trace/vfs/vfs_chmod| uint64|
/intel/osv/trace/vfs/vfs_chmod_ret| uint64|
/intel/osv/trace/vfs/vfs_chmod_err| uint64|
/intel/osv/trace/vfs/vfs_fchmod| uint64|
/intel/osv/trace/vfs/vfs_fchmod_ret| uint64|
/intel/osv/trace/vfs/vfs_fchown| uint64|
/intel/osv/trace/vfs/vfs_fchown_ret| uint64|

### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-osv/issues).

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-osv/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-osv/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
This is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

* Author: [Marcin Spoczynski](https://github.com/sandlbn/)

This software has been contributed by MIKELANGELO, a Horizon 2020 project co-funded by the European Union. https://www.mikelangelo-project.eu/
## Thank You
And **thank you!** Your contribution, through code and participation, is incredibly important to us.
