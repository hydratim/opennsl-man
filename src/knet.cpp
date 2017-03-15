#include <iostream>
#include <cstring>
#include <string>
#include <unordered_map>

#include <s11n.net/shellish/shellish.hpp> // eshell framework
#include <s11n.net/shellish/argv_parser.hpp>
#include <s11n.net/shellish/arguments.hpp>
#include <s11n.net/shellish/builtins.hpp>

#include <s11n.net/shellish/shellish_debuggering_macros.hpp> // CERR

#ifndef __common_
#define __common_
#include "common.cpp"
#endif

extern "C" {
#include "opennsl/error.h"
#include "opennsl/knet.h"
}


bool MODE_KNET = false;


opennsl_knet_netif_t get_netif (const shellish::arguments & args) {
    opennsl_knet_netif_t ret;
    opennsl_knet_netif_t_init(&ret);
    size_t i_len_args = args.argc();
    unsigned char* baseMac = mac_convert_str_to_bytes(args[i_len_args-1]);
    ret.type = OPENNSL_KNET_NETIF_T_TX_LOCAL_PORT;
    std::string tmp_port = args[i_len_args-2];
    ret.port = std::stoi(tmp_port);
    strcpy(ret.name, args[i_len_args-3].c_str());
    memcpy(ret.mac_addr, baseMac, 6);
    delete baseMac;
    return ret;
}

opennsl_knet_filter_t get_filter (const opennsl_knet_netif_t* netif) {
    opennsl_knet_filter_t ret;
    opennsl_knet_filter_t_init(&ret);
    ret.type = OPENNSL_KNET_FILTER_T_RX_PKT;
    ret.flags = OPENNSL_KNET_FILTER_F_STRIP_TAG;
    ret.dest_type = OPENNSL_KNET_DEST_T_NETIF;
    ret.dest_id = netif->id;
    ret.match_flags = OPENNSL_KNET_FILTER_M_INGPORT;
    ret.m_ingport = netif->port;
    return ret;
}

int KNETInit (const shellish::arguments & args) { 
    auto ret = opennsl_knet_init(0);
    if ( ret != OPENNSL_E_NONE ) {
        std::ostringstream err;
        err << "opennsl_knet_netif_init() failed " << opennsl_errmsg(ret);
        return 1;
    }
    return 0;
}

int KNETAdd (const shellish::arguments & args) { 
    auto netif = get_netif(args);
    auto ret = opennsl_knet_netif_create(0, &netif);
    if ( ret != OPENNSL_E_NONE ) {
        std::ostringstream err;
        shellish::ostream() << "opennsl_knet_netif_create() failed " << opennsl_errmsg(ret);
        return 1;
    }
    auto filter = get_filter(&netif);
    ret = opennsl_knet_filter_create(0, &filter);
    if ( ret != OPENNSL_E_NONE ) {
        std::ostringstream err;
        shellish::ostream() << "opennsl_knet_filter_create() failed " << opennsl_errmsg(ret);
        return 2;
    }
    shellish::ostream() << netif.id << std::endl;
    shellish::ostream() << filter.id << std::endl;
    return 0;
}

int KNETDelete (const shellish::arguments & args) {
    std::string tmp_id = args[args.argc()-2];
    std::string tmp_filter_id = args[args.argc()-1];
    int knet_id = std::stoi(tmp_id);
    int filter_id = std::stoi(tmp_filter_id);
    auto ret = opennsl_knet_netif_destroy(0, knet_id);
    if ( ret != OPENNSL_E_NONE ) {
        std::ostringstream err;
        shellish::ostream() << "opennsl_knet_netif_destroy() failed " << opennsl_errmsg(ret);
        return 1;
    }
    ret = opennsl_knet_filter_destroy(0, filter_id);
    if ( ret != OPENNSL_E_NONE ) {
        std::ostringstream err;
        shellish::ostream() << "opennsl_knet_netif_destroy() failed " << opennsl_errmsg(ret);
        return 2;
    }
    return 0;
}

int trav_fn(int unit, opennsl_knet_netif_t *info, void *user_data) {
    std::string mac;
    mac = mac_convert_bytes_to_sting(info->mac_addr);
    shellish::ostream() << info->id << " - " << info->mac_addr << " - " << info->port << " - " << info->vlan <<std::endl;
    return 0;
}
int trav_filter_fn(int unit, opennsl_knet_filter_t *info, void *user_data) {
    shellish::ostream() << info->id << " - " << info->m_ingport <<std::endl;
    return 0;
}

int KNETList (const shellish::arguments & args) { 
    void *n;
    auto ret = opennsl_knet_netif_traverse(0, trav_fn, n);
    if ( ret != OPENNSL_E_NONE ) {
        shellish::ostream() << "opennsl_knet_traverse() failed " << opennsl_errmsg(ret);
        return 1;
    }
    ret = opennsl_knet_filter_traverse(0, trav_filter_fn, n);
    if ( ret != OPENNSL_E_NONE ) {
        shellish::ostream() << "opennsl_knet_traverse() failed " << opennsl_errmsg(ret);
        return 1;
    }
    return 0;
}

int KNETExit (const shellish::arguments & args) { 
    MODE_KNET = false;
    shellish::ostream() << "KNET Mode: Deactivated" << std::endl;
    return 0; 
}

int KNET(const shellish::arguments & args) {
    std::unordered_map<std::string, shellish::command_handler_func> commands = { {"init", KNETInit}, {"add", KNETAdd}, {"delete", KNETDelete}, {"list", KNETList}, {"exit", KNETExit} };
    if (MODE_KNET) {
        auto search = commands.find(args[0]);
        if (search != commands.end()) {
            return search->second(args);
        } else {
            shellish::ostream() << "Command Error: \n User inputted: ";
            for( size_t i = 0; i < args.argc(); i++ ) {
                shellish::ostream() <<args[i] << " ";
            }
            shellish::ostream() << std::endl << std::endl << "Accepted commands are: " << std::endl;
            for (auto i = commands.begin(); i != commands.end(); ++i) {
                auto cur = i->first;
                shellish::ostream() << cur << std::endl;
            }
            return 0;
            }
    } else {
        if (args.argc() == 1){
            MODE_KNET = true;
            shellish::ostream() << "KNET Mode: Activated" << std::endl;
            return 0;
        } else {
            auto search = commands.find(args[1]);
            if (search != commands.end()) {
                return search->second(args);
            } else {
            shellish::ostream() << std::endl << std::endl << "Accepted commands are: " << std::endl;
            for (auto i = commands.begin(); i != commands.end(); ++i) {
                auto cur = i->first;
                shellish::ostream() << commands[cur] << std::endl;
            }
            return 0;
            }
        }
    }
}
