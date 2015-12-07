/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS"BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package osv

var virtioCounters = []string{"virtio_wait_for_queue", "virtio_enable_interrupts", "virtio_disable_interrupts",
	"virtio_kicked_event_idx", "virtio_add_buf", "virtio_net_rx_packet", "virtio_net_rx_wake", "virtio_net_fill_rx_ring",
	"virtio_net_fill_rx_ring_added", "virtio_net_tx_packet", "virtio_net_tx_failed_add_buf", "virtio_net_tx_no_space_calling_gc",
	"virtio_net_tx_packet_size", "virtio_net_tx_xmit_one_failed_to_post", "virtio_blk_read_config_capacity",
	"virtio_blk_read_config_size_max", "virtio_blk_read_config_seg_max", "virtio_blk_read_config_geometry",
	"virtio_blk_read_config_blk_size", "virtio_blk_read_config_topology", "virtio_blk_read_config_wce",
	"virtio_blk_read_config_ro", "virtio_blk_make_request_seg_max", "virtio_blk_make_request_readonly",
	"virtio_blk_wake", "virtio_blk_strategy", "virtio_blk_req_ok", "virtio_blk_req_unsupp", "virtio_blk_req_err"}

var netCounters = []string{"net_packet_in", "	net_packet_out", "net_packet_handling"}

var tcpCounters = []string{"tcp_state", "tcp_input_ack", "tcp_output", "tcp_output_error",
	"tcp_output_resched_start", "tcp_output_resched_end", "tcp_output_start", "tcp_output_ret",
	"tcp_output_just_ret", "tcp_output_cant_take_inp_lock", "tcp_timer_tso_flush",
	"tcp_timer_tso_flush_ret", "tcp_timer_tso_flush_err"}

var memoryCounters = []string{"memory_malloc", "memory_malloc_mempool", "memory_malloc_large", "memory_malloc_page",
	"memory_free", "memory_realloc", "memory_page_alloc", "memory_page_free", "memory_huge_failure", "memory_reclaim",
	"memory_wait", "memory_mmap", "memory_mmap_err", "memory_mmap_ret", "memory_munmap", "memory_munmap_err", "memory_munmap_ret"}

var calloutCounters = []string{"callout_init", "callout_reset", "callout_stop_wait", "callout_stop",
	"callout_thread_waiting", "callout_thread_dispatching"}

var waitCounters = []string{"waitqueue_wait", "waitqueue_wake_one", "waitqueue_wake_all"}

var asyncCounters = []string{"async_timer_task_create", "async_timer_task_destroy", "async_timer_task_reschedule",
	"async_timer_task_cancel", "async_timer_task_shutdown", "async_timer_task_fire", "async_timer_task_misfire",
	"async_timer_task_insert", "async_timer_task_remove", "async_worker_started", "async_worker_timer_fire",
	"async_worker_timer_fire_ret", "async_worker_fire", "async_worker_fire_ret"}

var vfsCounters = []string{"vfs_open", "vfs_open_ret", "vfs_open_err", "vfs_close",
	"vfs_close_ret", "vfs_close_err", "vfs_mknod", "vfs_mknod_ret", "vfs_mknod_err", "vfs_lseek",
	"vfs_lseek_ret", "vfs_lseek_err", "vfs_pread", "vfs_pread_ret", "vfs_pread_err", "vfs_pwrite",
	"vfs_pwrite_ret", "vfs_pwrite_err", "vfs_pwritev", "vfs_pwritev_ret", "vfs_pwritev_err", "vfs_ioctl",
	"vfs_ioctl_ret", "vfs_ioctl_err", "vfs_fsync", "vfs_fsync_ret", "vfs_fsync_err", "vfs_fstat",
	"vfs_fstat_ret", "vfs_fstat_err", "vfs_readdir", "vfs_readdir_ret", "vfs_readdir_err",
	"vfs_mkdir", "vfs_mkdir_ret", "vfs_mkdir_err", "vfs_rmdir", "vfs_rmdir_ret", "vfs_rmdir_err",
	"vfs_rename", "vfs_rename_ret", "vfs_rename_err", "vfs_chdir", "vfs_chdir_ret",
	"vfs_chdir_err", "vfs_fchdir", "vfs_fchdir_ret", "vfs_fchdir_err", "vfs_link",
	"vfs_link_ret", "vfs_link_err", "vfs_symlink", "vfs_symlink_ret", "vfs_symlink_err",
	"vfs_unlink", "vfs_unlink_ret", "vfs_unlink_err", "vfs_stat", "vfs_stat_ret", "vfs_stat_err",
	"vfs_lstat", "vfs_lstat_ret", "vfs_lstat_err", "vfs_statfs", "vfs_statfs_ret", "vfs_statfs_err",
	"vfs_fstatfs", "vfs_fstatfs_ret", "vfs_fstatfs_err", "vfs_getcwd", "vfs_getcwd_ret", "vfs_getcwd_err",
	"vfs_dup", "vfs_dup_ret", "vfs_dup_err", "vfs_dup3", "vfs_dup3_ret", "vfs_dup3_err", "vfs_fcntl",
	"vfs_fcntl_ret", "vfs_fcntl_err", "vfs_access", "vfs_access_ret", "vfs_access_err", "vfs_isatty",
	"vfs_isatty_ret", "vfs_isatty_err", "vfs_truncate", "vfs_truncate_ret", "vfs_truncate_err",
	"vfs_ftruncate", "vfs_ftruncate_ret", "vfs_ftruncate_err", "vfs_fallocate", "vfs_fallocate_ret",
	"vfs_fallocate_err", "vfs_utimes", "vfs_utimes_ret", "vfs_utimes_err", "vfs_utimensat",
	"vfs_utimensat_ret", "vfs_utimensat_err", "vfs_futimens", "vfs_futimens_ret", "vfs_futimens_err",
	"vfs_chmod", "vfs_chmod_ret", "vfs_chmod_err", "vfs_fchmod", "vfs_fchmod_ret", "vfs_fchown",
	"vfs_fchown_ret"}

var cpuMetrics = []string{"cputime"}
var memMetrics = []string{"free", "total"}
